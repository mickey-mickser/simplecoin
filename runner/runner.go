package runner

import (
	"fmt"
	"github.com/mickey-mickser/simplecoin"
	"github.com/mickey-mickser/simplecoin/pkg/repository"
	"github.com/superoo7/go-gecko/v3"
	"log"
	"time"
)

func ParseCoins(r *repository.PricesPostgres) error {
	cg := coingecko.NewClient(nil)

	var idCoins = map[string]string{
		"btc":  "bitcoin",
		"eth":  "ethereum",
		"ltc":  "litecoin",
		"doge": "dogecoin",
		"sol":  "solana",
		"usd":  "tether",
	}
	//var keyFound string
	//for key, value := range idCoins {
	//	if value == idCoins[key] {
	//		keyFound = key
	//		break
	//	}
	//}

	for {

		for _, value := range idCoins {
			price, err := cg.SimpleSinglePrice(value, "usd" /*id*/)
			if err != nil {
				log.Printf("Error getting price for %s: %v\n", value, err)
				continue
			}
			data := crypto.CryptoCoin{
				SymbolFrom: value,
				SymbolTo:   "usd",
				Price:      float64(price.MarketPrice),
			}

			if err = r.WriteCoinToDB(&data); err != nil {
				return fmt.Errorf("failed to write coin to DB for %s: %w", value, err)
			}
			time.Sleep(10 * time.Second)
		}
		time.Sleep(1 * time.Minute)
	}

	//}
	//time.Sleep(2 * time.Minute)

	return nil
}
