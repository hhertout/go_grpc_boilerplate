package interceptor

import (
	"google.golang.org/grpc"
)

// return all interceptor of the package
// They are executed in the order they are in the array
// The first interceptor in the array is the first to be executed
// The last interceptor in the array is the last to be executed
func GetInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		LoggingInterceptor,
		ApiKeyInterceptor,
	}
}
