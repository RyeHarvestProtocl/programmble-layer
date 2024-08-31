package programmableLayer

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/RyeHarvestProtocol/programmable-layer/evmClient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	UnimplementedGreeterServer
	UnimplementedFundTxHandlerServer
	UnimplementedMintTxHandlerServer
	UnimplementedClaimTxHandlerServer
	EvmClient *evmClient.EvmContractClient
	Config    *config.Config
}

func (s *Server) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + req.Name}, nil
}

// mint rye
func (s *Server) SubmitFundRequest(ctx context.Context, req *FundRequest) (*FundRequestReply, error) {
	txHash, err := s.EvmClient.SendMintRyeTransaction(req.UserAddress, s.Config.EvmClient.IndexerPrivateKey, req.RuneAmount)
	if err != nil {
		return &FundRequestReply{
			Success: false,
			Error:   err.Error(),
		}, err
	}
	fmt.Println("2222")

	return &FundRequestReply{
		Success: true,
		TxHash:  txHash,
		Error:   "",
	}, nil
}

// mint harvest
func (s *Server) SubmitMintRequest(ctx context.Context, req *MintRequest) (*MintRequestReply, error) {
	txHash, err := s.EvmClient.SendMintHarvestTransaction(req.UserAddress, s.Config.EvmClient.IndexerPrivateKey, req.MintAmount)
	if err != nil {
		return &MintRequestReply{
			Success: false,
			Error:   err.Error(),
			TxHash:  "",
		}, err
	}

	return &MintRequestReply{
		Success: true,
		TxHash:  txHash,
		Error:   "",
	}, nil
}

// claim harvest
func (s *Server) SubmitClaimRequest(ctx context.Context, req *ClaimRequest) (*ClaimRequestReply, error) {
	txHash, err := s.EvmClient.SendClaimHarvestTransaction(req.UserAddress, s.Config.EvmClient.IndexerPrivateKey, req.RoundId)
	if err != nil {
		return &ClaimRequestReply{
			Success: false,
			Error:   err.Error(),
			TxHash:  "",
		}, err
	}

	return &ClaimRequestReply{
		Success: true,
		TxHash:  txHash,
		Error:   "",
	}, nil
}

func StartServer(address string, config *config.Config) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	fmt.Println("this is config ")
	fmt.Println("this is config RpcUrl: ", config.EvmClient.RpcUrl)

	evmClient, err := evmClient.NewEvmContractClient(config)
	if err != nil {
		return err
	}

	fmt.Print("this is evmClient: ", evmClient)

	grpcServer := grpc.NewServer()
	RegisterGreeterServer(grpcServer, &Server{EvmClient: evmClient, Config: config})
	RegisterFundTxHandlerServer(grpcServer, &Server{EvmClient: evmClient, Config: config})
	RegisterMintTxHandlerServer(grpcServer, &Server{EvmClient: evmClient, Config: config})
	RegisterClaimTxHandlerServer(grpcServer, &Server{EvmClient: evmClient, Config: config})

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("gRPC server listening on %s", address)
	return grpcServer.Serve(lis)
}
