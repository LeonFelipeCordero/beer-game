package ports

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type IOrderApi interface {
	CreateOrder(ctx context.Context, receiverId string) (*model.Order, error)
	DeliverOrder(ctx context.Context, orderId string, amount int) (*model.Order, error)
}

type IOrderService interface {
	CreateOrder(ctx context.Context, receiverId string) (*domain.Order, error)
	DeliverOrder(ctx context.Context, orderId string, amount int) (*domain.Order, error)
	Get(ctx context.Context, orderId string) (*domain.Order, error)
}

type IOrderRepository interface {
	Save(ctx context.Context, order domain.Order) (*domain.Order, error)
	Get(ctx context.Context, orderId string) (*domain.Order, error)
}
