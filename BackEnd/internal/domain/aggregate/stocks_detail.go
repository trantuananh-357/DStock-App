package aggregate

import (
	"github.com/google/uuid"
)

type StocksByDetail struct {
	StockID     uuid.UUID `json:"stock_id"`
	StockSymbol string    `json:"stock_symbol"`
	CompanyName string    `json:"company_name"`
	Sector      string    `json:"sector"`
	MarketPrice uint64    `json:"market_price"`
}
