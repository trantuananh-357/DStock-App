package configuration

import (
	"dstock-backend/internal/common/composite"
	"dstock-backend/internal/common/httpresponse"
	"dstock-backend/internal/domain/usecase/stock"
	stockusecase "dstock-backend/internal/domain/usecase/stock"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockService struct {
	repo struct {
		stock *composite.StockRepoComposite
	}
	useCase struct {
		stock *composite.StocksUseCaseComposite
	}
	logger *zap.Logger
}

func NewStockService(stockRepoComposite *composite.StockRepoComposite, logger *zap.Logger) *StockService {
	return &StockService{
		logger: logger,
		repo: struct {
			stock *composite.StockRepoComposite
		}{
			stock: stockRepoComposite,
		},
		useCase: struct {
			stock *composite.StocksUseCaseComposite
		}{
			stock: &composite.StocksUseCaseComposite{
				GetAllStock: stockusecase.NewGetAllStockCase(stockRepoComposite.PersistentRepo, logger),
			},
		},
	}
}

func (s *StockService) StockService(ctx *gin.Context, req *stock.GetAllStockReq) *httpresponse.Response {
	res := &httpresponse.Response{}
	stockByDetail, _ := s.useCase.stock.GetAllStock.Execute(ctx, req)
	return res.TransformToSuccessOk(stockByDetail)
}
