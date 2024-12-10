package crypto

type CryptoCoin struct {
	SymbolFrom string  `json:"symbolFrom" db:"symbolFrom"`
	SymbolTo   string  `json:"symbolTo" db:"symbolTo"`
	Price      float64 `json:"price" db:"price"`
}
