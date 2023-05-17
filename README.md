<!-- ABOUT THE PROJECT -->
## About The Project
Microservices in Go with gRPC, API Gateway

### Built With
* golang
* grpc
* grpc-gateway
* gorm, gormmigrate
* viper

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites
```shell
go > 1.18
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Installation
* start
```shell
docker-compose up --build
```


* make
```shell
make generate_proto_user_management
make generate_proto_gateway
```