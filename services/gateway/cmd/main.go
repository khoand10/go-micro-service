package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-micro-service/services/gateway/config"
	pb "go-micro-service/services/gateway/protos"
	"go-micro-service/services/gateway/routers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	cfg := config.LoadConfig("./")
	log.Println(cfg)

	engine := gin.New()
	app := &routers.RestServer{
		Engine: engine,
		Config: cfg,
	}

	go func() {
		err := routers.Init(app)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	// Create a listener on TCP port
	//lis, err := net.Listen(config.DefaultGRPCServerConfig.Network, config.DefaultGRPCServerConfig.Address)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on connection ")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterGreeterHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    "localhost:8082",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := request.Name
	response := &pb.HelloResponse{
		Message: "Hello " + name,
	}
	return response, nil
}
