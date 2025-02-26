package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/whoami00911/gRPC-server/pkg/grpcPb"
	"github.com/whoami00911/gRPC-server/pkg/logger"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	logService grpcPb.LogServiceServer
	//grpcServerApi grpcPb.LogServiceServer
	addr   string
	port   string
	logger *logger.Logger
}

func NewGrpcServer(grpcLogServer grpcPb.LogServiceServer, logger *logger.Logger) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		logService: grpcLogServer,
		addr:       viper.GetString("server.addr"),
		port:       viper.GetString("server.port"),
		logger:     logger,
	}
}

func (s *Server) ListenAndServe() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.addr, s.port))
	//fmt.Printf("%s:%s", s.addr, s.port)
	if err != nil {
		s.logger.Error(fmt.Sprintf("ListenAndServe failed: %s", err))
		log.Panic("ListenAndServe failed")
	}

	grpcPb.RegisterLogServiceServer(s.grpcServer, s.logService)
	fmt.Println("gRPC server has been started")
	if err := s.grpcServer.Serve(listener); err != nil {
		s.logger.Error(fmt.Sprintf("ListenAndServe failed: %s", err))
		log.Panic("ListenAndServe failed")
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(done)
	}()
	select {
	case <-ctx.Done():
		s.logger.Error("Graceful shutdown timed out, forcing immediate stop")
		fmt.Println("Graceful shutdown timed out, forcing immediate stop")
		s.grpcServer.Stop()
		ctx.Err()
	case <-done:
		fmt.Println("Server shutdown gracefully")
		return nil
	}
	return nil
}
