package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"

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
				grpc_recovery.UnaryServerInterceptor(),
			)))
	v1.RegisterRedimoServiceServer(server, RedimoService{})
	reflection.Register(server)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		server.GracefulStop()
	}()
	log.Fatal(server.Serve(lis))
}

type RedimoService struct{}
