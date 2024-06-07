package service

import (
	"context"
	"github.com/Morozov-rondo/restapi-template/internal/repository"
	"github.com/Morozov-rondo/restapi-template/internal/service/service1"
	"github.com/Morozov-rondo/restapi-template/internal/service/service2"
)

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Service1: service1.NewService1(repos.Repository1),
		Service2: service2.NewService2(repos.Repository2),
	}
}

type Service struct {
	Service1
	Service2
}

type Service1 interface {
	Command1(ctx context.Context) error
	Command2(ctx context.Context) error
}

type Service2 interface {
	Command1(ctx context.Context) error
	Command2(ctx context.Context) error
}
