package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/hhertout/grpc_boilerplate/internal/interceptor"
	"github.com/hhertout/grpc_boilerplate/internal/server"
	"github.com/hhertout/grpc_boilerplate/pb"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// init logger
	logger, _ := zap.NewProduction()
	if os.Getenv("GO_ENV") == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	// Load .env file if not running in a docker container
	// else the env variables are set in the docker-compose file
	// default file is .env in both cases
	if os.Getenv("DOCKER_RUN") == "false" {
		err := godotenv.Load()
		if err != nil {
			logger.Fatal("Error loading .env file")
		}
	}

	// LOG warning if the server is running in development mode
	if os.Getenv("GO_ENV") == "development" {
		logger.Warn("⚠️ Caution : The server will be running under development mode 🔨🔨")
	}

	// Retrieve port from env variable
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Fatal("PORT env variable must be set and be an integer")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf((":%d"), port))
	if err != nil {
		logger.Fatal("Failed to create listener")
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.GetInterceptors(logger)...,
		),
	)
	reflection.Register(s)
	defer s.GracefulStop()

	pb.RegisterCalculatorServiceServer(s, &server.Server{})

	// #### Lauching server ####
	logger.Info("🚀 Server running", zap.Int("port", port))
	if err := s.Serve(listener); err != nil {
		s.GracefulStop()
		logger.Sugar().Fatalf("Fail to serve : %v", err)
	}
}
