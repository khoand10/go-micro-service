package handler

import (
	"context"
	"errors"
	pb "go-micro-service/services/user_management/protos"
	"go-micro-service/services/user_management/usecase"
)

func (s *Server) GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	userID, ok := ctx.Value("UserID").(uint64)
	if !ok {
		return nil, errors.New("invalid user_id")
	}
	in.UserId = userID
	return usecase.GetUserProfile(s.userRepository, in, s.cfg)
}
