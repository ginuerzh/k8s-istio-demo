package account

import (
	"context"
	"os"

	"github.com/ginuerzh/k8s-istio-demo/account-svc/api"
	userapi "github.com/ginuerzh/k8s-istio-demo/user-svc/api"
)

type Server struct {
	UserClient userapi.UserClient
}

func (s *Server) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	r, err := s.UserClient.Create(ctx, &userapi.UserCreateRequest{
		User: &userapi.UserDetail{
			Username: req.Username,
		},
	})
	if err != nil {
		return nil, err
	}

	host, _ := os.Hostname()
	return &api.RegisterResponse{Id: r.Id, Host: host + "/" + r.Name}, nil
}

func (s *Server) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	r, err := s.UserClient.Detail(ctx, &userapi.UserDetailRequest{
		Id: req.Username,
	})
	if err != nil {
		return nil, err
	}

	host, _ := os.Hostname()
	return &api.LoginResponse{Id: r.Id, Host: host + "/" + r.Name}, nil
}

func (s *Server) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error) {
	r, err := s.UserClient.Delete(ctx, &userapi.UserDeleteRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	host, _ := os.Hostname()
	return &api.LogoutResponse{
		Id:   r.Id,
		Host: host + "/" + r.Name,
	}, nil
}
