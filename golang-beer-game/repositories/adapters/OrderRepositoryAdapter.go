package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
)

type OrderRepositoryAdapter struct {
}

func NewOrderRepository() ports.IOrderRepository {
	return &OrderRepositoryAdapter{}
}

func (o OrderRepositoryAdapter) Save(ctx context.Context, order domain.Order) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepositoryAdapter) Get(ctx context.Context, orderId string) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepositoryAdapter) DeleteAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
