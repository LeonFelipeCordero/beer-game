package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBoardPlayerAndOrder(t *testing.T) {
	neo4j.ConfigureDatabase()
	ctx := context.Background()
	repository := neo4j.NewRepository()
	boardRepository := NewBoardRepository(repository)
	playerRepository := NewPlayerRepository(repository, boardRepository)
	orderRepository := NewOrderRepository(repository, playerRepository)
	t.Run("create all relations between boards and player easily", func(t *testing.T) {
		board := createBoard(ctx, boardRepository)
		retailer := createPlayer(ctx, playerRepository, board, "RETAILER")
		wholesaler := createPlayer(ctx, playerRepository, board, "WHOLESALER")

		order := domain.Order{
			Amount:         5,
			OriginalAmount: 5,
			Status:         domain.StatusPending,
			OrderType:      domain.OrderTypePlayerOrder,
			CreatedAt:      time.Now().UTC(),
			Receiver:       retailer.Id,
			Sender:         wholesaler.Id,
		}
		savedOrder, err := orderRepository.Save(ctx, order)
		assert.Nil(t, err, "error when saving order")
		validateOrder(ctx, t, orderRepository, *savedOrder)
		deliverOrder(ctx, t, orderRepository, *savedOrder)
	})
}

func validateOrder(ctx context.Context, t *testing.T, repository ports.IOrderRepository, order domain.Order) {
	savedOrder, err := repository.Get(ctx, order.Id)
	assert.Nil(t, err, "error when saving order")
	assert.Equal(t, savedOrder.Amount, order.Amount, "wrong amount")
	assert.Equal(t, savedOrder.OriginalAmount, order.OriginalAmount, "wrong original amount")
	assert.Equal(t, savedOrder.Status, order.Status, "wrong status")
	assert.Equal(t, savedOrder.OrderType, order.OrderType, "wrong order type")
	assert.Equal(t, savedOrder.CreatedAt, order.CreatedAt, "wrong creation time")
	assert.Equal(t, savedOrder.Receiver, order.Receiver, "wrong receiver")
	assert.Equal(t, savedOrder.Sender, order.Sender, "wrong sender")
}

func deliverOrder(ctx context.Context, t *testing.T, repository ports.IOrderRepository, order domain.Order) {
	savedOrder, _ := repository.Get(ctx, order.Id)
	savedOrder.Status = domain.StatusDelivered
	savedOrder.Amount = 1
	savedOrder, err := repository.Save(ctx, *savedOrder)
	assert.Nil(t, err, "error when saving order")
	assert.Equal(t, savedOrder.Amount, 1, "wrong amount")
	assert.Equal(t, savedOrder.Status, domain.StatusDelivered, "wrong status")
}
