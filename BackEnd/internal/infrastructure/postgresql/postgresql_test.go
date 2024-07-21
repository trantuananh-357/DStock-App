package postgresql

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func TestConnectToDB(t *testing.T) {
	db, err := gorm.Open(postgres.Open("postgresql://neondb_owner:Z4zf7CFvglQK@ep-shrill-sound-a5g0r3g2.us-east-2.aws.neon.tech/dstock?sslmode=require"), &gorm.Config{})
	if err != nil {
		t.Fatal("Error connecting to the database: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Log("failed to get generic database object: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		t.Log("failed to ping database: %w", err)
	}
	t.Log("Successfully connected to the database")
}
