package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("RPC failed with status: %v", st.Code())
	}

	log.Printf("gRPC method=%s; time taken=%v", info.FullMethod, time.Since(start))
	return resp, err
}
