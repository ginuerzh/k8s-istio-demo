//go:generate protoc -I../api --go_out=plugins=grpc:../api ../api/account.proto
//go:generate protoc -I../api --swagger_out=logtostderr=true:../api ../api/account.proto

package main

import (
	"context"
	"log"
	"net"
	"os"

	api "github.com/ginuerzh/k8s-istio-demo/account-svc/api"
	"github.com/ginuerzh/k8s-istio-demo/account-svc/svc/account"
	"github.com/ginuerzh/k8s-istio-demo/account-svc/svc/health"
	authapi "github.com/ginuerzh/k8s-istio-demo/auth-svc/api"
	userapi "github.com/ginuerzh/k8s-istio-demo/user-svc/api"
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
	log.Println("listening for account grpc service on", addr)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			unaryServerRecoveryInterceptor(),
			unaryServerForwardTraceHeadersInterceptor(traceHeaders),
			// unaryServerOpenTracingInterceptor(tracer),
			// unaryServerAuthInterceptor(),
			unaryServerLoggingInterceptor(),
		),
	)

	userSvc := os.Getenv("USER_SERVICE")
	if userSvc == "" {
		userSvc = "user:8000"
	}
	userConn, err := grpc.DialContext(context.Background(), userSvc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial user service: %v", err)
	}

	authSvc := os.Getenv("AUTH_SERVICE")
	if authSvc == "" {
		authSvc = "auth:8000"
	}
	authConn, err := grpc.DialContext(context.Background(), authSvc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial auth service: %v", err)
	}

	api.RegisterAccountServer(s, &account.Server{
		UserClient: userapi.NewUserClient(userConn),
		AuthClient: authapi.NewAuthClient(authConn),
	})
	grpc_health_v1.RegisterHealthServer(s, &health.Server{})
	reflection.Register(s)

	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
