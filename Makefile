generate_proto_gateway:
	protoc -I ./internal/grpc_protos/user_management \
      --go_out ./services/gateway/protos/user_management --go_opt paths=source_relative \
      --go-grpc_out ./services/gateway/protos/user_management --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./services/gateway/protos/user_management --grpc-gateway_opt paths=source_relative \
      ./internal/grpc_protos/user_management/user_management.proto

generate_proto_user_management:
	protoc -I ./internal/grpc_protos/user_management \
		--go_out=./services/user_management/protos --go_opt=paths=source_relative \
        --go-grpc_out=./services/user_management/protos --go-grpc_opt=paths=source_relative \
        ./internal/grpc_protos/user_management/user_management.proto