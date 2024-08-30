package server

import (
	"context"
	"fmt"

	"github.com/hhertout/grpc_boilerplate/internal/service"
	"github.com/hhertout/grpc_boilerplate/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCalculatorServer
}

func (s *Server) Add(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	fmt.Println("I'm in the add function")
	if in.A == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Can't set 0 to A")
	}

	return &pb.CalculationResponse{
		Result: int64(service.Add(int(in.A), int(in.B))),
	}, nil
}
