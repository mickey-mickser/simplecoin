package usecase //service
import (
	"context"

	"github.com/mickey-mickser/simplecoin"
	"github.com/mickey-mickser/simplecoin/pkg/repository"
)

type CryptoList interface {
	GetAll(ctx context.Context, extraColumns []string) ([]crypto.CryptoCoin, error)
	GetCoin(ctx context.Context, coin crypto.CryptoCoin) (float64, error)
}
type UseCase struct {
	CryptoList
}

func NewPricesUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		CryptoList: NewRatesUseCase(repo.CryptoList),
	}
}
