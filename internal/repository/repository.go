package repository

import (
	"context"
	"gRPC-Server/pkg/grpcPb"
	"gRPC-Server/pkg/logger"
)

type Logging interface {
	Insert(ctx context.Context, req grpcPb.LogItem)
}

type Repository struct {
	Logging
}

func NewRepo(repo Logging, logger *logger.Logger) *Repository {
	return &Repository{
		Logging: repo,
	}
}
func (r *Repository) Insert(ctx context.Context, req *grpcPb.LogRequest) {

}
