package main

import (
	"fmt"
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/handler"
	"go-micro-service/services/user_management/infra/mysql"
	"go-micro-service/services/user_management/middlaware"
	"go-micro-service/services/user_management/migration/schema"
	"go-micro-service/services/user_management/migration/seed_data"
	pb "go-micro-service/services/user_management/protos"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//TODO fake wait mysql start, implement wait for mysql started
	time.Sleep(5 * time.Second)

	cfg := config.LoadConfig("./")

	db, err := mysql.Connect(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	err = schema.Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	err = seed_data.Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	userRepository := mysql.NewUserRepositoryImpl(db)
	server := handler.CreateServer(userRepository, cfg)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middlaware.JWTAUth(cfg),
		))
	pb.RegisterUserManagementServer(grpcServer, server)
	log.Printf("[INFO] start grpc server listening %d", cfg.UserManagementPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.UserManagementPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
