package main

import (
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	v1 "github.com/dbProjectRED/redimo/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":4772")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger, _ := zap.NewProduction()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpczap.UnaryServerInterceptor(logger),
			)))
	v1.RegisterRedimoServiceServer(server, RedimoService{})
	log.Fatal(server.Serve(lis))
}

type RedimoService struct{}
