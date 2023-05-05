package handler

import (
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/domain/repository"
	pb "go-micro-service/services/user_management/protos"
)

type Server struct {
	pb.UnimplementedUserManagementServer
	userRepository repository.UserRepository
	cfg            *config.Config
}

func CreateServer(userRepository repository.UserRepository, cfg *config.Config) *Server {
	return &Server{
		userRepository: userRepository,
		cfg:            cfg,
	}
}
