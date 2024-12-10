package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mickey-mickser/simplecoin"
	"github.com/sirupsen/logrus"
	"strings"
)

type PricesPostgres struct {
	db *sqlx.DB
}

func NewPricesPostgres(db *sqlx.DB) *PricesPostgres {
	return &PricesPostgres{db: db}
}

func (p *PricesPostgres) WriteCoinToDB(crypto *crypto.CryptoCoin) error {
	logrus.Printf("Inserting/Updating coin: %s -> %s = %f\n", crypto.SymbolFrom, crypto.SymbolTo, crypto.Price)
	//query := `INSERT INTO crypto_prices(symbolFrom, symbolTo, price)
	//		  VALUES($1, $2, $3)
	//		  ON CONFLICT(symbolFrom, symbolTo)
	//		  DO UPDATE SET price = EXCLUDED.price`
	//

	query := "INSERT INTO crypto_prices(symbolFrom, symbolTo, price) VALUES($1, $2, $3) "
	conflictQuery := "ON CONFLICT(symbolFrom, symbolTo) DO UPDATE SET price = EXCLUDED.price"
	finalQuery := query + conflictQuery
	_, err := p.db.Exec(finalQuery, crypto.SymbolFrom, crypto.SymbolTo, crypto.Price)
	if err != nil {
		logrus.Println("error writing to db:", err)
		return err
	}
	return err
}

func (p *PricesPostgres) GetAll(ctx context.Context, extraColumns []string) ([]crypto.CryptoCoin, error) {
	//query := "SELECT symbolFrom, symbolTo, price FROM crypto_prices"
	//rows, err := p.db.QueryContext(ctx, query)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//var coins []crypto.CryptoCoin
	//for rows.Next() {
	//	var coin crypto.CryptoCoin
	//	if err := rows.Scan(&coin.SymbolFrom, &coin.SymbolTo, &coin.Price); err != nil {
	//		return nil, err
	//	}
	//	coins = append(coins, coin)
	//}
	//
	//return coins, nil
	columns := []string{"symbolFrom", "symbolTo", "price"}
	for _, col := range extraColumns {
		columns = append(columns, col)
	}
	query := "SELECT " + strings.Join(columns, ", ") + " FROM crypto_prices"
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	///////
	var coins []crypto.CryptoCoin
	for rows.Next() {
		var coin crypto.CryptoCoin
		if err = rows.Scan(&coin.SymbolFrom, &coin.SymbolTo, &coin.Price); err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}
	return coins, nil
}

func (p *PricesPostgres) GetCoin(ctx context.Context, coin crypto.CryptoCoin) (float64, error) {
	//var coinPrice float64
	//query := fmt.Sprintf("SELECT price FROM crypto_prices WHERE symbolFrom=$1")
	//row := p.db.QueryRowContext(ctx, query, coin.SymbolFrom)
	//if err := row.Scan(&coinPrice); err != nil {
	//	return 0, err
	//}
	//return coinPrice, nil
	var coinPrice float64
	query := "SELECT price FROM crypto_prices WHERE"
	args := []interface{}{}
	if coin.SymbolFrom != "" {
		query += " AND symbolFrom = $1"
		args = append(args, coin.SymbolFrom)
	}
	err := p.db.QueryRowContext(ctx, query, args...).Scan(&coinPrice)
	if err != nil {
		return 0, err
	}
	return coinPrice, nil
}
