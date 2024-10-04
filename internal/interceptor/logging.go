package interceptor

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ContextKey string

// @Global middleware
//
// LoggingInterceptor logs the time taken for the request to be processed
// and the method that was called
func LoggingInterceptor(l *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		// Add logger to context
		loggerKey := ContextKey("logger")
		ctx = context.WithValue(ctx, loggerKey, l)

		start := time.Now()
		resp, err := handler(ctx, req)
		if err != nil {
			st, _ := status.FromError(err)
			l.Error("RPC failed", zap.String("status", st.Code().String()))
		}

		md, _ := metadata.FromIncomingContext(ctx)
		userAgent := "unknown"
		if userAgentHeader := md.Get("user-agent"); len(userAgentHeader) > 0 {
			userAgent = userAgentHeader[0]
		}

		st, _ := status.FromError(err)
		statusCode := st.Code().String()

		logFields := []zap.Field{
			zap.String("method", info.FullMethod),
			zap.String("user_agent", userAgent),
			zap.Duration("time_taken", time.Since(start)),
			zap.String("status", statusCode),
		}

		if err != nil {
			l.Error("RPC failed", logFields...)
		} else {
			l.Info("RPC call", logFields...)
		}

		return resp, err
	}
}
