package auth

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ginuerzh/k8s-istio-demo/auth-svc/api"
	"github.com/go-redis/redis/v7"
)

type Server struct {
	PrivKey     *rsa.PrivateKey
	PubKey      *rsa.PublicKey
	TTL         time.Duration
	RedisClient *redis.Client
}

func (s *Server) CreateToken(ctx context.Context, req *api.CreateTokenRequest) (*api.CreateTokenResponse, error) {
	exp := time.Now().Add(s.TTL).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid":  req.Uid,
		"exp":  exp,
		"plat": req.Platform,
	})
	t, err := token.SignedString(s.PrivKey)
	if err != nil {
		return nil, err
	}

	if _, err := s.RedisClient.
		Set(fmt.Sprintf("auth:token:uid:%s", req.Uid), t, s.TTL).
		Result(); err != nil {
		return nil, err
	}

	return &api.CreateTokenResponse{
		Token:      t,
		Expiration: exp,
	}, nil
}

func (s *Server) decode(t string) (map[string]interface{}, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return s.PubKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token invalid")
	}

	return claims, nil
}

func (s *Server) DecodeToken(ctx context.Context, req *api.DecodeTokenRequest) (*api.DecodeTokenResponse, error) {
	claims, err := s.decode(req.Token)
	if err != nil {
		return nil, err
	}

	resp := &api.DecodeTokenResponse{}
	if v, ok := claims["uid"].(string); ok {
		resp.Uid = v
	}
	if v, ok := claims["exp"].(float64); ok {
		resp.Expiration = int64(v)
	}
	if v, ok := claims["plat"].(string); ok {
		resp.Platform = v
	}

	if token, _ := s.RedisClient.
		Get(fmt.Sprintf("auth:token:uid:%s", resp.Uid)).
		Result(); token != req.Token {
		return nil, errors.New("invalid token")
	}
	return resp, nil
}

func (s *Server) DeleteToken(ctx context.Context, req *api.DeleteTokenRequest) (*api.DeleteTokenResponse, error) {
	claims, err := s.decode(req.Token)
	if err != nil {
		return nil, err
	}

	var uid string
	if v, ok := claims["uid"].(string); ok {
		uid = v
	}
	s.RedisClient.Del(fmt.Sprintf("auth:token:uid:%s", uid))

	return &api.DeleteTokenResponse{Uid: uid}, nil
}
