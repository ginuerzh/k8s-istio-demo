//go:generate protoc -I../api --go_out=plugins=grpc:../api ../api/user.proto
//go:generate protoc -I../api --swagger_out=logtostderr=true:../api ../api/user.proto

package main

import (
	"context"
	"log"
	"net"
	"os"

	api "github.com/ginuerzh/k8s-istio-demo/user-svc/api"
	"github.com/ginuerzh/k8s-istio-demo/user-svc/svc/health"
	"github.com/ginuerzh/k8s-istio-demo/user-svc/svc/user"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	log.Println("listening for user grpc service on", addr)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			unaryServerRecoveryInterceptor(),
			// unaryServerOpenTracingInterceptor(tracer),
			// unaryServerAuthInterceptor(),
			unaryServerLoggingInterceptor(),
		),
	)

	mongoAddr := os.Getenv("MONGO_URI")
	if mongoAddr == "" {
		mongoAddr = "mongodb://mongo:27017"
	}

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoAddr))
	if err != nil {
		log.Fatalf("failed to init mongodb: %v", err)
	}

	api.RegisterUserServer(s, &user.Server{
		MongoClient: mongoClient,
	})
	grpc_health_v1.RegisterHealthServer(s, &health.Server{})
	reflection.Register(s)

	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
