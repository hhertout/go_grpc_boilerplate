package interceptor

import (
	"google.golang.org/grpc"
)

func GetInterceptor() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		LoggingInterceptor,
		ApiKeyInterceptor,
	}
}
