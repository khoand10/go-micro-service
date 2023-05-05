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
```

### Installation
* start
```shell
go mod tidy
docker-compose up -d
```


* make
```shell
make generate_proto_user_management
make generate_proto_gateway
```