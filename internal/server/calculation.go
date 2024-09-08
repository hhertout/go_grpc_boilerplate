package server

import (
	"context"

	"github.com/hhertout/grpc_boilerplate/internal/service"
	"github.com/hhertout/grpc_boilerplate/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Result: service.Add(in.A, in.B),
	}, nil
}

func (s *Server) Divide(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	if in.B == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Can't divide by 0 !")
	}

	return &pb.AddResponse{
		Result: service.Divide(in.A, in.B),
	}, nil
}
