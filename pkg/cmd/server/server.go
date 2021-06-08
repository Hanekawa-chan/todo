package server

import (
	"context"
	"github.com/Hanekawa-chan/todo/pkg/protocol/grpc"
	v1 "github.com/Hanekawa-chan/todo/pkg/service/v1"
)

type Config struct {
}

func RunServer() error {
	ctx := context.Background()

	db := v1.Db{}
	v1API := v1.NewToDoServiceServer(&db)

	return grpc.RunServer(ctx, v1API)
}
