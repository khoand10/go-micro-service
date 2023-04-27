package handler

import (
	"context"
	pb "go-micro-service/services/user_management/protos"
)

func (s *Server) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
