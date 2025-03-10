package repositorymq

import (
	"context"

	"github.com/whoami00911/gRPC-server/internal/repository"
	"github.com/whoami00911/gRPC-server/pkg/grpcPb"
	"github.com/whoami00911/gRPC-server/pkg/logger"
)

type LoggingMq interface {
	Insert(ctx context.Context, req *grpcPb.LogItem) error
}

type RepositoryMq struct {
	mongo  *repository.Mongo
	logger *logger.Logger
}

func NewRepoMq(repo *repository.Mongo, logger *logger.Logger) *RepositoryMq {
	return &RepositoryMq{
		mongo:  repo,
		logger: logger,
	}
}

func (r *RepositoryMq) Insert(ctx context.Context, req *grpcPb.LogItem) error {
	return r.mongo.Insert(ctx, req)
}
