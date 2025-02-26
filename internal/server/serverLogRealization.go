package server

import (
	"context"
	"fmt"
	"gRPC-Server/pkg/grpcPb"
	"gRPC-Server/pkg/logger"
)

type Logging interface {
	Insert(ctx context.Context, req *grpcPb.LogRequest) error
}

type LogServer struct {
	grpcPb.UnimplementedLogServiceServer
	Logging
	logger *logger.Logger
}

func NewLogServer(logserver Logging, logger *logger.Logger) *LogServer {
	return &LogServer{
		Logging: logserver,
		logger:  logger,
	}
}

func (l *LogServer) Log(ctx context.Context, req *grpcPb.LogRequest) (*grpcPb.Emty, error) {
	if err := l.Logging.Insert(ctx, req); err != nil {
		l.logger.Error(fmt.Sprintf("Insert Error: %s", err))
		return &grpcPb.Emty{}, err
	}
	return &grpcPb.Emty{}, nil
}
