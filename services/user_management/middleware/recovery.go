package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				log.Println(err)
			}
		}()

		return handler(ctx, req)
	}
}
