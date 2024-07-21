package postgresql

import (
	"dstock-backend/internal/domain/aggregate"
	"dstock-backend/internal/domain/models"
	"dstock-backend/internal/domain/repo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type postgresqlStockRepoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgresqlStockRepo(db *gorm.DB, logger *zap.Logger) repo.StocksRepo[*gorm.DB] {
	return &postgresqlStockRepoImpl{
		db:     db,
		logger: logger,
	}
}

// GetByAllStock implements repo.StocksRepo.
func (r *postgresqlStockRepoImpl) GetByAllStock(ctx *gin.Context, query *repo.Query) ([]*aggregate.StocksByDetail, error) {
	var stocks []*aggregate.StocksByDetail
	builder := r.db.WithContext(ctx).Select(models.STOCKTABLE_ID, models.STOCKTABLE_SYMBOL, models.STOCKTABLE_COMPANY_NAME, models.STOCKTABLE_SECTOR, models.STOCKTABLE_MARKET_PRICE).Limit(query.Limit).
		Offset(query.Offset).Table(models.STOCKTABLE).Find(&stocks)
	if query.SortBy != "" {
		builder.Order(query.SortBy + " " + query.Order)
	} else {
		builder.Order("stocks.id DESC ")
	}
	return stocks, nil
}

// GetByStockId implements repo.StocksRepo.
func (r *postgresqlStockRepoImpl) GetByStockId(*gin.Context, uuid.UUID, *repo.Query) ([]*aggregate.StocksByDetail, error) {
	panic("unimplemented")
}
