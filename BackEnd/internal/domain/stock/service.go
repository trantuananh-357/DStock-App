package domain

import (
	"github.com/labstack/echo"
)

type Service interface {
	GetStockByUUID(ctx echo.Context, uuid string) *Stock
	GetAllStock(ctx echo.Context, limit, offset int) []*Stock
}

type service struct {
	storage Storage
}

func newService(storage Storage) Service {
	return &service{storage: storage}
}

// GetAllStock implements Service.
func (s *service) GetAllStock(ctx echo.Context, limit int, offset int) []*Stock {
	return s.storage.GetAll(limit, offset)
}

func (s *service) GetStockByUUID(ctx echo.Context, uuid string) *Stock {
	return s.storage.GetOne(uuid)
}
