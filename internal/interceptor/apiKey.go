package interceptor

import (
	"context"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// @Global middleware
//
// ApiKeyInterceptor checks if the API key provided in the header is valid
// Checked with the API_KEY environment variable
// It returns an error if the API key is invalid
func ApiKeyInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	logger := ctx.Value(ContextKey("logger")).(*zap.Logger)

	apiKey := os.Getenv("API_KEY")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Error("No metadata found")
		return nil, status.Errorf(codes.PermissionDenied, "No metadata found")
	}

	headerApiKey := md.Get("x-api-key")
	if len(headerApiKey) == 0 {
		logger.Error("No API key found in the header", zap.String("API_KEY", ""))
		return nil, status.Errorf(codes.PermissionDenied, "INVALID API KEY")
	}

	if headerApiKey[0] != apiKey {
		logger.Error("Invalid API key", zap.String("API_KEY", headerApiKey[0]))
		return nil, status.Errorf(codes.PermissionDenied, "INVALID API KEY")
	}

	resp, err := handler(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		logger.Sugar().Errorf("RPC failed with status: %v", st.Code())
	}
	return resp, err
}
