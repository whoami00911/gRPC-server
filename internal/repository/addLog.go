package repository

import (
	"context"
	"fmt"
	"gRPC-Server/pkg/grpcPb"
	"gRPC-Server/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	db     *mongo.Database
	logger *logger.Logger
}

func NewDatabaseInicialize(db *mongo.Database, logger *logger.Logger) *Mongo {
	return &Mongo{
		db:     db,
		logger: logger,
	}
}

func (m *Mongo) Insert(ctx context.Context, req grpcPb.LogItem) error {
	_, err := m.db.Collection("logs").InsertOne(ctx, req)
	if err != nil {
		m.logger.Error(fmt.Sprintf("Error Insert to mongo: %s", err))
		return err
	}
	return nil
}
