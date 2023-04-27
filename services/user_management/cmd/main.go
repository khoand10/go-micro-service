package main

import (
	"fmt"
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/handler"
	pb "go-micro-service/services/user_management/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig("./")
	log.Println(cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.UserManagementPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := handler.CreateServer()

	pb.RegisterUserManagementServer(grpcServer, server)
	log.Printf("[INFO] start http server listening %d", cfg.UserManagementPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
