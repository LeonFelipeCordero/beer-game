package adapters

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/google/uuid"
)

type OrderRepositoryAdapterFaker struct {
	orders map[string]domain.Order
}

func NewOrderRepositoryFaker() ports.IOrderRepository {
	return &OrderRepositoryAdapterFaker{
		orders: make(map[string]domain.Order),
	}
}

func (o OrderRepositoryAdapterFaker) Save(ctx context.Context, order domain.Order) (*domain.Order, error) {
	if order.Id == "" {
		id, _ := uuid.NewUUID()
		order.Id = id.String()
		o.orders[id.String()] = order
		return &order, nil
	}

	_, err := o.Get(ctx, order.Id)
	if err == nil {
		delete(o.orders, order.Id)
		o.orders[order.Id] = order
		return &order, nil
	}

	return nil, err
}

func (o OrderRepositoryAdapterFaker) Get(ctx context.Context, orderId string) (*domain.Order, error) {
	for key, value := range o.orders {
		if key == orderId {
			return &value, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Order with id %s doesn't exit", orderId))
}

func (o OrderRepositoryAdapterFaker) DeleteAll(ctx context.Context) {
	for key, _ := range o.orders {
		delete(o.orders, key)
	}
}
