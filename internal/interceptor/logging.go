package interceptor

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type ContextKey string

// @Global middleware
//
// LoggingInterceptor logs the time taken for the request to be processed
// and the method that was called
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Add logger to context
	loggerKey := ContextKey("logger")
	ctx = context.WithValue(ctx, loggerKey, logger)

	start := time.Now()
	resp, err := handler(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		logger.Sugar().Errorf("RPC failed with status: %v", st.Code())
	}

	logger.Info("RPC call", zap.String("method", info.FullMethod), zap.Duration("time_taken", time.Since(start)))

	return resp, err
}
