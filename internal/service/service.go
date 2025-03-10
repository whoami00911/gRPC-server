package service

import (
	"context"

	"github.com/whoami00911/gRPC-server/pkg/grpcPb"
	"github.com/whoami00911/gRPC-server/pkg/logger"
)

type Logging interface {
	Insert(ctx context.Context, req *grpcPb.LogRequest) error
}

type Service struct {
	Logging
	logger *logger.Logger
}

func NewService(repo Logging, logger *logger.Logger) *Service {
	return &Service{
		Logging: repo,
		logger:  logger,
	}
}

func (s *Service) Insert(ctx context.Context, req *grpcPb.LogRequest) error {
	return s.Logging.Insert(ctx, req)
}
