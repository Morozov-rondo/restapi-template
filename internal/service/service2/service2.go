package service2

import (
	"context"
	"github.com/Morozov-rondo/restapi-template/internal/repository"
)

type Service2 struct {
	repo repository.Repository2
}

func NewService2(repo repository.Repository2) *Service2 {
	return &Service2{
		repo: repo,
	}
}

func (s *Service2) Command1(ctx context.Context) error {

	// TODO: implement
	return nil
}

func (s *Service2) Command2(ctx context.Context) error {

	// TODO: implement
	return nil
}
