package user

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ginuerzh/k8s-istio-demo/user-svc/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	MongoClient *mongo.Client
}

func (s *Server) Create(ctx context.Context, req *api.UserCreateRequest) (*api.UserDetail, error) {
	col := s.MongoClient.
		Database("istio-demo").
		Collection("users")
	err := col.FindOne(ctx, bson.M{"username": req.GetUser().Username}).Err()
	if err == nil {
		return nil, errors.New("user exists")
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	user := &api.UserDetail{
		Username: req.GetUser().Username,
		Name:     req.GetUser().Name,
		Age:      req.GetUser().Age,
		Avatar:   req.GetUser().Avatar,
	}

	res, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = fmt.Sprintf("%v", res.InsertedID)

	return user, nil
}

func (s *Server) Detail(ctx context.Context, req *api.UserDetailRequest) (*api.UserDetail, error) {
	user := &api.UserDetail{}
	col := s.MongoClient.
		Database("istio-demo").
		Collection("users")
	err := col.FindOne(ctx, bson.M{"username": req.Id}).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) List(ctx context.Context, req *api.UserListRequest) (*api.UserListResponse, error) {
	resp := &api.UserListResponse{}

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
	user := &api.UserDetail{
		Username: req.Id,
	}
	col := s.MongoClient.
		Database("istio-demo").
		Collection("users")
	_, err := col.DeleteOne(ctx, bson.M{"username": req.Id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
