package handler

import (
	"context"
	pb "go-micro-service/services/user_management/protos"
	"go-micro-service/services/user_management/usecase"
)

func (s *Server) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {
	return usecase.Login(s.userRepository, input, s.cfg)
}

func (s *Server) Register(ctx context.Context, input *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return usecase.Register(s.userRepository, input, s.cfg)
}

func (s *Server) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return usecase.RefreshToken(s.userRepository, in, s.cfg)
}
