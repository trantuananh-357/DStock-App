package domain

import "github.com/google/uuid"

type Stock struct {
	StockID      uuid.UUID `json:"stock_id"`
	TickerSymbol string    `json:"ticker_symbol"`
	CompanyName  string    `json:"company_name"`
	Sector       string    `json:"sector"`
	Industry     string    `json:"industry"`
}
