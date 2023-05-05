package handler

import (
	"context"
	"fmt"
	pb "go-micro-service/services/user_management/protos"
	"go-micro-service/services/user_management/usecase"
)

func (s *Server) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("login handle ==============")
	//response := &pb.LoginResponse{
	//	Token:  "h2i34u2yi3ghe278fy49uijse" + input.Email,
	//	UserId: 1,
	//	Email:  input.Email,
	//}
	//s.userRepository.GetByEmail(input.Email)
	//return response, nil
	return usecase.Login(s.userRepository, input, s.cfg)
}
