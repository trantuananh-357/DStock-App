package models

import (
	"time"

	"github.com/google/uuid"
)

var STOCKTABLE = "stocks"

var (
	STOCKTABLE_ID           = STOCKTABLE + ".id"
	STOCKTABLE_SYMBOL       = STOCKTABLE + ".symbol"
	STOCKTABLE_COMPANY_NAME = STOCKTABLE + ".company_name"
	STOCKTABLE_SECTOR       = STOCKTABLE + ".sector"
	STOCKTABLE_MARKET_PRICE = STOCKTABLE + ".market_price"
	STOCKTABLE_CREATED_AT   = STOCKTABLE + ".created_at"
	STOCKTABLE_UPDATED_AT   = STOCKTABLE + ".updated_at"
)

type Stocks struct {
	StockID     uuid.UUID `gorm:"column:id;primaryKey;not null"`
	StockSymbol string    `gorm:"column:symbol"`
	CompanyName string    `gorm:"column:company_name"`
	Sector      string    `gorm:"column:sector"`
	MarketPrice uint64    `gorm:"column:market_price"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
