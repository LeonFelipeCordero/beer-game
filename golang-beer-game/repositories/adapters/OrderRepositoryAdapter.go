package adapters

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j/entities"
	"strconv"
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

// Save todo query players with one call
func (o OrderRepositoryAdapter) Save(ctx context.Context, order domain.Order) (*domain.Order, error) {
	receiver, err := o.playerRepository.Get(ctx, order.Receiver)
	if err != nil {
		return nil, err
	}
	receiverNode := &entities.PlayerNode{}
	receiverNode.FromPlayer(*receiver)
	sender, err := o.playerRepository.Get(ctx, order.Sender)
	if err != nil {
		return nil, err
	}
	senderNode := &entities.PlayerNode{}
	senderNode.FromPlayer(*sender)

	orderNode := &entities.OrderNode{}
	orderNode.FromOrder(order)
	orderNode.Receiver = receiverNode
	orderNode.Sender = senderNode
	err = o.repository.SaveDepth(ctx, orderNode)
	if err != nil {
		return nil, err
	}

	savedOrder := orderNode.ToOrder()
	return &savedOrder, nil
}

func (o OrderRepositoryAdapter) Get(ctx context.Context, orderId string) (*domain.Order, error) {
	entityId, _ := strconv.ParseInt(orderId, 0, 64)
	orderNode := &entities.OrderNode{}
	err := o.repository.LoadDepth(ctx, entityId, orderNode)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", orderId),
			err,
		)
	}

	savedOrder := orderNode.ToOrder()
	return &savedOrder, nil
}

func (o OrderRepositoryAdapter) DeleteAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
