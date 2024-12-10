package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mickey-mickser/simplecoin"
)

type CryptoList interface {
	GetAll(ctx context.Context, extraColumns []string) ([]crypto.CryptoCoin, error)
	GetCoin(ctx context.Context, coin crypto.CryptoCoin) (float64, error)
}
type Repository struct {
	CryptoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CryptoList: NewPricesPostgres(db),
	}
}
