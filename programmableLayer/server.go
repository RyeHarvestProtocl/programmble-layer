package programmableLayer

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	UnimplementedGreeterServer
	UnimplementedFundTxHandlerServer
	UnimplementedMintTxHandlerServer
	UnimplementedClaimTxHandlerServer
}

func (s *Server) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + req.Name}, nil
}

func (s *Server) SubmitFundRequest(ctx context.Context, req *FundRequest) (*FundRequestReply, error) {
	return &FundRequestReply{
		Success: true,
		Error:   "",
	}, nil
}

func (s *Server) SubmitMintRequest(ctx context.Context, req *MintRequest) (*MintRequestReply, error) {
	return &MintRequestReply{
		Success: true,
		Error:   "",
		TrackId: 1,
	}, nil
}

func (s *Server) SubmitClaimRequest(ctx context.Context, req *ClaimRequest) (*ClaimRequestReply, error) {
	return &ClaimRequestReply{
		Success: true,
		Error:   "",
	}, nil
}

// StartServer initializes and starts the gRPC server.
func StartServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterGreeterServer(grpcServer, &Server{})
	RegisterFundTxHandlerServer(grpcServer, &Server{}) // Add this line
	RegisterMintTxHandlerServer(grpcServer, &Server{})
	RegisterClaimTxHandlerServer(grpcServer, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("gRPC server listening on %s", address)
	return grpcServer.Serve(lis)
}
