package service1

import (
	"context"
	"github.com/Morozov-rondo/restapi-template/internal/repository"
)

type Service1 struct {
	repo repository.Repository1
}

func NewService1(repo repository.Repository1) *Service1 {
	return &Service1{
		repo: repo,
	}
}

func (s *Service1) Command1(ctx context.Context) error {

	// TODO: implement
	return nil
}

func (s *Service1) Command2(ctx context.Context) error {

	// TODO: implement
	return nil
}
