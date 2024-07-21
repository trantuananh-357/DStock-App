package composite

import (
	"dstock-backend/internal/domain/repo"
	stock_usecase "dstock-backend/internal/domain/usecase/stock"

	"gorm.io/gorm"
)

type StockRepoComposite struct {
	PersistentRepo repo.StocksRepo[*gorm.DB]
}
type StocksUseCaseComposite struct {
	GetAllStock stock_usecase.GetAllStock[*gorm.DB]
}
