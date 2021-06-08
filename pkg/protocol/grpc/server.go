package grpc

import (
	"context"
	"github.com/Hanekawa-chan/todo/pkg/api/v1"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC grpc...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC grpc...")
	return server.Serve(listen)
}