package rabbitmqservice

import (
	"context"

	repositorymq "github.com/whoami00911/gRPC-server/internal/repository/rabbitmq"
	"github.com/whoami00911/gRPC-server/pkg/grpcPb"
	"github.com/whoami00911/gRPC-server/pkg/logger"
)

type LoggingMq interface {
	Insert(ctx context.Context, req *grpcPb.LogItem) error
}

type ServiceMq struct {
	logger *logger.Logger
	repo   *repositorymq.RepositoryMq
}

func NewServiceMq(repo *repositorymq.RepositoryMq, logger *logger.Logger) *ServiceMq {
	return &ServiceMq{
		repo:   repo,
		logger: logger,
	}
}

func (s *ServiceMq) Insert(ctx context.Context, req *grpcPb.LogItem) error {
	return s.repo.Insert(ctx, req)
}
