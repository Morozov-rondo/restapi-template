package repository2

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func NewDB2(db *sqlx.DB) *DB2 {

	return &DB2{
		db: db,
	}
}

type DB2 struct {
	db *sqlx.DB
}

func (p *DB2) Command1(ctx context.Context) error {

	// TODO: implement
	return nil
}
func (p *DB2) Command2(ctx context.Context) error {

	// TODO: implement
	return nil
}
