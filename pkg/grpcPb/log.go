package grpcPb

import (
	"errors"
	"time"

	"github.com/whoami00911/gRPC-server/pkg/logger"
)

const (
	ENTITY_USER   = "USER"
	ENTITY_ENTITY = "ENTITY"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

var (
	entities = map[string]LogRequest_Entities{
		ENTITY_USER:   LogRequest_USER,
		ENTITY_ENTITY: LogRequest_ENTITY,
	}

	actions = map[string]LogRequest_Actions{
		ACTION_CREATE:   LogRequest_CREATE,
		ACTION_UPDATE:   LogRequest_UPDATE,
		ACTION_GET:      LogRequest_GET,
		ACTION_DELETE:   LogRequest_DELETE,
		ACTION_REGISTER: LogRequest_REGISTER,
		ACTION_LOGIN:    LogRequest_LOGIN,
	}
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	UserID    int64     `bson:"user_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (LogRequest_Entities, error) {
	logger := logger.GetLogger()
	val, ok := entities[entity]
	if !ok {
		logger.Error("Cant convert string entity to protobuf")
		return 0, errors.New("cant convert string entity to protobuf")
	}
	return val, nil
}

func ToPbAction(action string) (LogRequest_Actions, error) {
	logger := logger.GetLogger()
	val, ok := actions[action]
	if !ok {
		logger.Error("Cant convert string entity to protobuf")
		return 0, errors.New("cant convert string entity to protobuf")
	}
	return val, nil
}
