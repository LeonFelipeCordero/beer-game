package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
)

type OrderRepositoryAdapter struct {
	repository       neo4j.IRepository
	playerRepository ports.IPlayerRepository
}

func NewOrderRepository(
	repository neo4j.IRepository,
	playerRepository ports.IPlayerRepository,
) ports.IOrderRepository {
	return &OrderRepositoryAdapter{
		repository:       repository,
		playerRepository: playerRepository,
	}
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
