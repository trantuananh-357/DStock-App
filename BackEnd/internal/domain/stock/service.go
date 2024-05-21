package stock

import "context"

type Service interface {
	GetStockByUUID(ctx context.Context, uuid string) *Stock
	GetAllStock(ctx context.Context, limit, offset int) []*Stock
}

type service struct {
	storage Storage
}

// GetAllStock implements Service.
func (s *service) GetAllStock(ctx context.Context, limit int, offset int) []*Stock {
	return s.storage.GetAll(limit, offset)
}

func newService(storage Storage) Service {
	return &service{storage: storage}
}
func (s *service) GetStockByUUID(ctx context.Context, uuid string) *Stock {
	return s.storage.GetOne(uuid)
}