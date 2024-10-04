package interceptor

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// return all interceptor of the package
// They are executed in the order they are in the array
// The first interceptor in the array is the first to be executed
// The last interceptor in the array is the last to be executed
func GetInterceptors(logger *zap.Logger) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		LoggingInterceptor(logger),
		ApiKeyInterceptor,
	}
}
