package service

import (
	"context"
	"fmt"
	"gRPC-Server/pkg/grpcPb"
	"gRPC-Server/pkg/logger"
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
	if err := s.Logging.Insert(ctx, req); err != nil {
		s.logger.Error(fmt.Sprintf("Error insert data: %s", err))
		return err
	}
	return nil
}
