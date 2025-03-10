package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/whoami00911/gRPC-server/internal/repository"
	repositorymq "github.com/whoami00911/gRPC-server/internal/repository/rabbitmq"
	rabbitmqserver "github.com/whoami00911/gRPC-server/internal/server/rabbitmq-server"
	rabbitmqservice "github.com/whoami00911/gRPC-server/internal/service/rabbitmq"
	"github.com/whoami00911/gRPC-server/pkg/logger"
)

func init() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.ReadInConfig()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка при чтении конфигурации: %v", err)
	}
}

func main() {
	logger := logger.GetLogger()
	repoMq := repositorymq.NewRepoMq(repository.NewDatabaseInicialize(repository.NewMongoConnect(), logger), logger)
	serviceMq := rabbitmqservice.NewServiceMq(repoMq, logger)
	serverMq, err := rabbitmqserver.NewRabbitMQServer(serviceMq, logger)
	if err != nil {
		logger.Errorf("serverMq error: %s", err)
		log.Fatalf("serverMq error: %s", err)
	}
	fmt.Println("Rabbitmq started listening")
	serverMq.ListenAndServeMq()
}
