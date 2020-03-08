package main

import (
	"context"
	"log"
	"time"

	"github.com/ginuerzh/k8s-istio-demo/user-svc/api"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	conn, err := grpc.Dial("192.168.100.100:80",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithAuthority("grpc.istio.demo"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := api.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.Detail(ctx, &api.UserDetailRequest{Id: "123"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
}
