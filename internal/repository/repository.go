package repository

import (
	"context"
	"fmt"

	"github.com/whoami00911/gRPC-server/pkg/grpcPb"
	"github.com/whoami00911/gRPC-server/pkg/logger"
)

type Logging interface {
	Insert(ctx context.Context, req grpcPb.LogItem) error
}

type Repository struct {
	Logging
	logger *logger.Logger
}

func NewRepo(repo Logging, logger *logger.Logger) *Repository {
	return &Repository{
		Logging: repo,
		logger:  logger,
	}
}
func (r *Repository) Insert(ctx context.Context, req *grpcPb.LogRequest) error {
	logItem := grpcPb.LogItem{
		Entity:    req.GetEntity().String(),
		Action:    req.GetAction().String(),
		EntityID:  req.GetEntityId(),
		UserID:    req.GetUserId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}
	if err := r.Logging.Insert(ctx, logItem); err != nil {
		r.logger.Error(fmt.Sprintf("Error insert logItem: %s", err))
		return err
	}
	return nil
}
