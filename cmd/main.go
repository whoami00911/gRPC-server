package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"github.com/whoami00911/gRPC-server/internal/repository"
	"github.com/whoami00911/gRPC-server/internal/server"
	"github.com/whoami00911/gRPC-server/internal/service"
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
	repo := repository.NewRepo(repository.NewDatabaseInicialize(repository.NewMongoConnect(), logger), logger)
	service := service.NewService(repo, logger)
	logServer := server.NewLogServer(service, logger)
	grpcServer := server.NewGrpcServer(logServer, logger)
	go grpcServer.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := grpcServer.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Shutdown error: %s", err))
	}
}
