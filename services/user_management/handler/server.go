package handler

import (
	pb "go-micro-service/services/user_management/protos"
)

type Server struct {
	pb.UnimplementedUserManagementServer
}

func CreateServer() *Server {
	return &Server{}
}
