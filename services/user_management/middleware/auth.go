package middleware

import (
	"context"
	"go-micro-service/pkg/array"
	"go-micro-service/services/user_management/config"
	"go-micro-service/services/user_management/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var PublicFunctions = []string{
	"/UserManagement/Login",
	"/UserManagement/Register",
	"/UserManagement/RefreshToken",
}

func JWTAUth(cfg *config.Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if array.Contains(PublicFunctions, info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(401, "token invalid")
		}

		token := md.Get("Authorization")
		tokenStr := token[0]

		claims, err := jwt.VerifyToken(tokenStr, cfg.JwtSecretKey)
		if err != nil {
			return nil, status.Errorf(401, "token invalid")
		}
		newContext := context.WithValue(ctx, "UserID", claims.UserId)
		return handler(newContext, req)
	}
}
