package server

import (
	"context"
	"gRPC-Server/pkg/grpcPb"
)

type Logging interface {
	Insert(ctx context.Context, req *grpcPb.LogRequest)
}

type LogServer struct {
	grpcPb.UnimplementedLogServiceServer
	Logging
}

func NewLogServer(logserver Logging) *LogServer {
	return &LogServer{
		Logging: logserver,
	}
}

func (l *LogServer) Log(ctx context.Context, req *grpcPb.LogRequest) (*grpcPb.Emty, error) {

	return &grpcPb.Emty{}, nil
}
