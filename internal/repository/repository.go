package repository

import (
	"context"
	"github.com/Morozov-rondo/restapi-template/internal/repository/repository1"
	"github.com/Morozov-rondo/restapi-template/internal/repository/repository2"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Repository1: repository1.NewDB1(db),
		Repository2: repository2.NewDB2(db),
	}
}

type Repository struct {
	Repository1
	Repository2
}

type Repository1 interface {
	Command1(ctx context.Context) error
	Command2(ctx context.Context) error
}

type Repository2 interface {
	Command1(ctx context.Context) error
	Command2(ctx context.Context) error
}
