package ports

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type IOrderApi interface {
	CreateOrder(ctx context.Context, receiverId string) (*model.Order, error)
	DeliverOrder(ctx context.Context, orderId string, amount int) (*model.Response, error)
	LoadByBoard(ctx context.Context, boardId string) ([]*model.Order, error)
	LoadByPlayer(ctx context.Context, playerId string) ([]*model.Order, error)
}

type IOrderService interface {
	CreateOrder(ctx context.Context, receiverId string) (*domain.Order, error)
	DeliverOrder(ctx context.Context, orderId string, amount int) (*domain.Order, error)
	Get(ctx context.Context, orderId string) (*domain.Order, error)
	LoadByBoard(ctx context.Context, boardId string) ([]*domain.Order, error)
	LoadByPlayer(ctx context.Context, playerId string) ([]*domain.Order, error)
}

type IOrderRepository interface {
	Save(ctx context.Context, order domain.Order) (*domain.Order, error)
	Get(ctx context.Context, orderId string) (*domain.Order, error)
	LoadByBoard(ctx context.Context, boardId string) ([]*domain.Order, error)
	LoadByPlayer(ctx context.Context, playerId string) ([]*domain.Order, error)
	DeleteAll(ctx context.Context)
}
