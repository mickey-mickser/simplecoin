package usecase

import (
	"context"
	"github.com/mickey-mickser/simplecoin"
	"github.com/mickey-mickser/simplecoin/pkg/repository"
)

type CoinsUseCase struct {
	repo repository.CryptoList
}

func NewRatesUseCase(repo repository.CryptoList) *CoinsUseCase {
	return &CoinsUseCase{repo: repo}
}

func (u *CoinsUseCase) GetAll(ctx context.Context, extraColumns []string) ([]crypto.CryptoCoin, error) {
	return u.repo.GetAll(ctx, nil)
}
func (u *CoinsUseCase) GetCoin(ctx context.Context, coin crypto.CryptoCoin) (float64, error) {
	return u.repo.GetCoin(ctx, coin)
}
