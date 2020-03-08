//go:generate protoc -I../api --go_out=plugins=grpc:../api ../api/auth.proto
//go:generate protoc -I../api --swagger_out=logtostderr=true:../api ../api/auth.proto

package main

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	api "github.com/ginuerzh/k8s-istio-demo/auth-svc/api"
	"github.com/ginuerzh/k8s-istio-demo/auth-svc/svc/auth"
	"github.com/ginuerzh/k8s-istio-demo/auth-svc/svc/health"
	"github.com/go-redis/redis/v7"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := os.Getenv("GRPC_ADDR")
	if addr == "" {
		addr = ":8000"
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening for auth grpc service on", addr)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			unaryServerRecoveryInterceptor(),
			// unaryServerOpenTracingInterceptor(tracer),
			// unaryServerAuthInterceptor(),
			unaryServerLoggingInterceptor(),
		),
	)

	auths, err := initAuthServer()
	if err != nil {
		log.Fatalf("failed to init auth service: %v", err)
	}
	api.RegisterAuthServer(s, auths)
	grpc_health_v1.RegisterHealthServer(s, &health.Server{})
	reflection.Register(s)

	errChan := make(chan error)
	go func() {
		if err := s.Serve(ln); err != nil {
			errChan <- err
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	defer func() {
		log.Println("graceful stop server...")
		s.GracefulStop()
		log.Println("graceful stop server... done!")
	}()

	select {
	case <-sigs:
	case err := <-errChan:
		log.Println("serve err:", err)
	}
}

func initAuthServer() (*auth.Server, error) {
	privKey, pubKey, err := loadRSAKey()
	if err != nil {
		return nil, err
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	cli := redis.NewClient(opt)
	//	if _, err := cli.Ping().Result(); err != nil {
	//		log.Fatalf("redis ping: %v", err)
	//	}

	ttl, _ := time.ParseDuration(os.Getenv("TTL"))
	if ttl <= 0 {
		ttl = 3600 * time.Second
	}

	s := &auth.Server{
		PrivKey:     privKey,
		PubKey:      pubKey,
		RedisClient: cli,
		TTL:         ttl,
	}

	return s, nil
}

func loadRSAKey() (priv *rsa.PrivateKey, pub *rsa.PublicKey, err error) {
	ss, err := ioutil.ReadFile("/ssl/privkey.pem")
	if err != nil {
		return
	}
	priv, err = jwt.ParseRSAPrivateKeyFromPEM(ss)
	if err != nil {
		return
	}

	ss, err = ioutil.ReadFile("/ssl/pubkey.pem")
	if err != nil {
		return
	}
	pub, err = jwt.ParseRSAPublicKeyFromPEM(ss)
	return
}
