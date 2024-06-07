package repository1

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func NewDB1(db *sqlx.DB) *DB1 {

	return &DB1{
		db: db,
	}
}

type DB1 struct {
	db *sqlx.DB
}

func (d *DB1) Command1(ctx context.Context) error {

	// TODO: implement
	return nil
}
func (d *DB1) Command2(ctx context.Context) error {

	// TODO: implement
	return nil
}
