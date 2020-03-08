package user

import (
	"context"
	"fmt"
	"os"

	"github.com/ginuerzh/k8s-istio-demo/user-svc/api"
)

type Server struct{}

func (s *Server) Create(ctx context.Context, req *api.UserCreateRequest) (*api.UserDetail, error) {
	host, _ := os.Hostname()

	user := &api.UserDetail{
		Id:       req.GetUser().Username,
		Username: req.GetUser().Username,
		Name:     host,
		Age:      req.GetUser().Age,
		Avatar:   req.GetUser().Avatar,
	}

	return user, nil
}

func (s *Server) Detail(ctx context.Context, req *api.UserDetailRequest) (*api.UserDetail, error) {
	host, _ := os.Hostname()

	user := &api.UserDetail{
		Id:     req.GetId(),
		Name:   host,
		Age:    18,
		Avatar: "http://example.com/avatar.jpg",
	}
	return user, nil
}

func (s *Server) List(ctx context.Context, req *api.UserListRequest) (*api.UserListResponse, error) {
	host, _ := os.Hostname()
	resp := &api.UserListResponse{}

	for i := 0; i < 3; i++ {
		resp.Users = append(resp.Users, &api.UserDetail{
			Id:     fmt.Sprintf("user%d", i),
			Name:   host,
			Age:    int32(10 + i),
			Avatar: fmt.Sprintf("http://www.example.com/avatar%d.jpg", i),
		})
	}
	return resp, nil
}

func (s *Server) Update(ctx context.Context, req *api.UserUpdateRequest) (*api.UserDetail, error) {
	host, _ := os.Hostname()

	user := req.User
	if user == nil {
		user = &api.UserDetail{}
	}
	user.Id = req.Id
	user.Name = host

	return user, nil
}

func (s *Server) Delete(ctx context.Context, req *api.UserDeleteRequest) (*api.UserDetail, error) {
	host, _ := os.Hostname()

	user := &api.UserDetail{
		Id:   req.GetId(),
		Name: host,
	}

	return user, nil
}
