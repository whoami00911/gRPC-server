package repository

import (
	"context"
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

func (m *Mongo) Insert(_ context.Context, req grpcPb.LogItem) {

}
