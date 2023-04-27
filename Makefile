generate_proto_user_management:
	protoc --go_out=./services/account_management/ \
        --go-grpc_out=./services/user_management/protos \
    	--go-grpc_opt=paths=source_relative \
        --proto_path=./internal/grpc-protos/user_management.proto