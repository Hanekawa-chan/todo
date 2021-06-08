package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/Hanekawa-chan/todo/pkg/protocol/grpc"
	v1 "github.com/Hanekawa-chan/todo/pkg/service/v1"
)

type Config struct {
	GRPCPort string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	db := v1.Db{}
	v1API := v1.NewToDoServiceServer(&db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
