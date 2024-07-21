package repo

import (
	"dstock-backend/internal/domain/aggregate"

	"github.com/gin-gonic/gin"
)

type Query struct {
	SortBy string `form:"sort_by"`
	Order  string `form:"order_by"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type StocksRepo[StocksType any] interface {
	// Create(*gin.Context, *models.Stocks, StocksType) error
	// Update(*gin.Context, *models.Stocks, StocksType) error
	// Delete(*gin.Context, *models.Stocks, StocksType) error
	// GetByStockId(*gin.Context, uuid.UUID, *Query) ([]*aggregate.StocksByDetail, error)
	GetByAllStock(*gin.Context, *Query) ([]*aggregate.StocksByDetail, error)
}
