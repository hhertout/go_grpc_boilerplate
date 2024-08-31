package interceptor

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ApiKeyInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	apiKey := os.Getenv("API_KEY")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "No metadata found")
	}

	headerApiKey := md.Get("x-api-key")[0]

	if headerApiKey != apiKey {
		return nil, status.Errorf(codes.PermissionDenied, "INVALID API KEY")
	}

	resp, err := handler(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("RPC failed with status: %v", st.Code())
	}
	return resp, err
}
