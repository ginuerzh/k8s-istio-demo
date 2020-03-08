package main

import (
	"context"
	"encoding/base64"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func unaryServerLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		log.Println("info:", info.FullMethod, info.Server)
		log.Println("incoming md:", md)

		md, _ = metadata.FromOutgoingContext(ctx)
		log.Println("outgoing md:", md)

		log.Println("req:", req)

		resp, err = handler(ctx, req)

		log.Println("resp:", resp)
		return
	}
}

func unaryServerAuthInterceptor() grpc.UnaryServerInterceptor {
	authFunc := func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "basic")
		if err != nil {
			return nil, err
		}
		data, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}
		if string(data) != "foo:bar" {
			return nil, grpc.Errorf(codes.Unauthenticated, "authorization failed")
		}
		return ctx, nil
	}
	return grpc_auth.UnaryServerInterceptor(authFunc)
}

func unaryServerRecoveryInterceptor() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor()
}

func unaryServerOpenTracingInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer))
}

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

func unaryServerForwardTraceHeadersInterceptor(eaders []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var kv []string
		md, _ := metadata.FromIncomingContext(ctx)
		if md.Len() > 0 {
			for _, k := range traceHeaders {
				if v := md.Get(k); len(v) > 0 {
					kv = append(kv, k, v[0])
				}
			}
		}
		if len(kv) > 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, kv...)
		}

		return handler(ctx, req)
	}
}
