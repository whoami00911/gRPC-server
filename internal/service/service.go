package service

import (
	"context"
	"gRPC-Server/pkg/grpcPb"
)

type Logging interface {
	Insert(ctx context.Context, req *grpcPb.LogRequest)
}

type Service struct {
	Logging
}

func NewService(repo Logging) *Service {
	return &Service{
		Logging: repo,
	}
}

func (s *Service) Insert(_ context.Context, req *grpcPb.LogRequest) {}
