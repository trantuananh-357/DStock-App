package stock

import (
	"dstock-backend/internal/domain/aggregate"
	"dstock-backend/internal/domain/repo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetAllStockReq struct {
	Query *repo.Query
}

type GetAllStock[TxType any] interface {
	Execute(ctx *gin.Context, req *GetAllStockReq) ([]*aggregate.StocksByDetail, error)
}

type defaultGetAllStockCase[TxType any] struct {
	stockRepo repo.StocksRepo[TxType]
	logger    *zap.Logger
}

func NewGetAllStockCase[TxType any](stockRepo repo.StocksRepo[TxType], logger *zap.Logger) GetAllStock[TxType] {
	return &defaultGetAllStockCase[TxType]{
		stockRepo: stockRepo,
		logger:    logger,
	}
}

func (s *defaultGetAllStockCase[TxType]) Execute(ctx *gin.Context, req *GetAllStockReq) ([]*aggregate.StocksByDetail, error) {
	return s.stockRepo.GetByAllStock(ctx, req.Query)
}
