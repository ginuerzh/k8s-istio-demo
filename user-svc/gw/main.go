//go:generate protoc -I../api --grpc-gateway_out=logtostderr=true:../api ../api/user.proto

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	api "github.com/ginuerzh/k8s-istio-demo/user-svc/api"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	// trace header to propagate.
	traceHeaders = []string{
		"x-ot-span-context",
		"x-request-id",

		// Zipkin headers
		"b3",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",

		// Jaeger header (for native client)
		"uber-trace-id",
	}
)

func incomingHeaderMatcherFunc(key string) (string, bool) {
	k := strings.ToLower(key)
	for _, v := range traceHeaders {
		if v == k {
			return key, true
		}
	}
	return runtime.DefaultHeaderMatcher(key)
}

func run() error {
	grpcEndpoint := os.Getenv("GRPC_ENDPOINT")
	if grpcEndpoint == "" {
		grpcEndpoint = ":8000"
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(incomingHeaderMatcherFunc))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterUserHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	gwAddr := os.Getenv("GW_ADDR")
	if gwAddr == "" {
		gwAddr = ":8080"
	}
	log.Printf("listening on %s forward to %s", gwAddr, grpcEndpoint)

	return http.ListenAndServe(gwAddr, mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
