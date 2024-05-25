package domain

type Storage interface {
	GetOne(uuid string) *Stock
	GetAll(limit, offset int) []*Stock
	Create(stock *Stock) (*Stock, error)
	Delete(uuid string) error
}

