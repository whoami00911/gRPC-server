package main

import (
	"context"
	"fmt"
	"gRPC-Server/internal/repository"
	"gRPC-Server/internal/server"
	"gRPC-Server/internal/service"
	"gRPC-Server/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
