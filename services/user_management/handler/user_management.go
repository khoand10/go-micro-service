package handler

import (
	"context"
	"errors"
	pb "go-micro-service/services/user_management/protos"
	"go-micro-service/services/user_management/usecase"
)

func (s *Server) GetUserProfile(ctx context.Context, input *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	userID, ok := ctx.Value("UserID").(uint64)
	if !ok {
		return nil, errors.New("invalid user_id")
	}
	input.UserId = userID
	return usecase.GetUserProfile(s.userRepository, input)
}

func (s *Server) GetUser(ctx context.Context, input *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	request := usecase.GetUserRequest{
		UserID: input.UserId,
	}
	user, err := usecase.GetUserDetail(s.userRepository, request)
	if err != nil {
		return nil, err
	}

	res := &pb.GetUserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}
	return res, nil
}
