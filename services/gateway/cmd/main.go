package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-micro-service/services/gateway/config"
	protos "go-micro-service/services/gateway/protos/user_management"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig("./")

	/*
		engine := gin.New()
		app := &routers.RestServer{
			Engine: engine,
			Config: cfg,
		}

		go func() {
		err := routers.Init(app)
		if err != nil {
			log.Fatalln(err)
			return
		}
		}()
	*/

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	httpServerAddress := fmt.Sprintf(":%d", cfg.GatewayRestPort)
	err := createUserManagementClient(ctx, mux, cfg)
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to User Management service: %v", err)
	}

	log.Printf("[INFO] start http server listening %d", cfg.GatewayRestPort)
	err = http.ListenAndServe(httpServerAddress, mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUserManagementClient(ctx context.Context, mux *runtime.ServeMux, cfg *config.Config) error {
	userManagementGrpcEndpoint := cfg.UserManagementURI
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := protos.RegisterUserManagementHandlerFromEndpoint(ctx, mux, userManagementGrpcEndpoint, opts)
	if err != nil {
		return err
	}
	return nil
}
