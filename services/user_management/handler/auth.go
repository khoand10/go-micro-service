package handler

import (
	"context"
	"github.com/jinzhu/copier"
	pb "go-micro-service/services/user_management/protos"
	"go-micro-service/services/user_management/usecase"
)

func (s *Server) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {
	var request usecase.LoginRequest
	err := copier.Copy(&request, input)
	if err != nil {
		return nil, err
	}

	loginRes, err := usecase.Login(s.userRepository, &request, s.cfg)
	if err != nil {
		return nil, err
	}
	res := &pb.LoginResponse{
		Token:        loginRes.Token,
		RefreshToken: loginRes.RefreshToken,
	}
	return res, nil
}

func (s *Server) Register(ctx context.Context, input *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var request usecase.RegisterRequest
	err := copier.Copy(&request, input)
	if err != nil {
		return nil, err
	}
	newUser, err := usecase.Register(s.userRepository, &request)
	if err != nil {
		return nil, err
	}
	res := &pb.RegisterResponse{
		Id:    newUser.Id,
		Email: newUser.Email,
	}
	return res, nil
}

func (s *Server) RefreshToken(ctx context.Context, input *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	var request usecase.RefreshTokenRequest
	err := copier.Copy(&request, input)
	if err != nil {
		return nil, err
	}
	refreshTokenRes, err := usecase.RefreshToken(s.userRepository, &request, s.cfg)
	if err != nil {
		return nil, err
	}
	res := &pb.RefreshTokenResponse{
		Token:        refreshTokenRes.Token,
		RefreshToken: refreshTokenRes.RefreshToken,
	}
	return res, nil
}
