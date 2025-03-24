package adapters

import (
	"context"
	testingutil "github.com/LeonFelipeCordero/golang-beer-game/testing"
	"testing"
	"time"

	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/stretchr/testify/assert"
)

func TestBoardPlayerAndOrder(t *testing.T) {
	ctx := context.Background()
	queries := testingutil.SetupDatabaseConnection(ctx)
	boardRepository := NewBoardRepository(queries)
	playerRepository := NewPlayerRepository(queries)
	orderRepository := NewOrderRepository(queries)

	testingutil.Clean(ctx, queries)

	t.Run("create all relations between boards and player easily", func(t *testing.T) {
		board := createBoard(ctx, boardRepository)
		retailer := createPlayer(ctx, playerRepository, board, "RETAILER")
		wholesaler := createPlayer(ctx, playerRepository, board, "WHOLESALER")

		var ids []string
		savedOrder := createOrder(ctx, t, orderRepository, *retailer, *wholesaler)
		validateOrder(ctx, t, orderRepository, *savedOrder)
		deliverOrder(ctx, t, orderRepository, *savedOrder)

		ids = append(ids, savedOrder.Id)
		for i := 0; i < 4; i++ {
			extraOrder := createOrder(ctx, t, orderRepository, *retailer, *wholesaler)
			ids = append(ids, extraOrder.Id)
		}
		loadOrdersByBoard(ctx, t, orderRepository, board.Id, ids)
		loadOrdersByPlayer(ctx, t, orderRepository, retailer.Id, ids)
	})
}

func createOrder(ctx context.Context, t *testing.T, repository OrderRepositoryAdapter, receiver domain.Player, sender domain.Player) *domain.Order {
	order := domain.Order{
		Amount:         5,
		OriginalAmount: 5,
		Status:         domain.StatusPending,
		OrderType:      domain.OrderTypePlayerOrder,
		CreatedAt:      time.Now().UTC(),
		Receiver:       receiver.Id,
		Sender:         sender.Id,
	}
	savedOrder, err := repository.Save(ctx, order)
	assert.Nil(t, err, "error when saving order")
	return savedOrder
}

func validateOrder(ctx context.Context, t *testing.T, repository OrderRepositoryAdapter, order domain.Order) {
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

func deliverOrder(ctx context.Context, t *testing.T, repository OrderRepositoryAdapter, order domain.Order) {
	savedOrder, err := repository.MarkAsFilled(ctx, order.Id, int64(1))
	assert.Nil(t, err, "error when saving order")
	fetchedOrder, err := repository.Get(ctx, order.Id)
	assert.Nil(t, err, "error when saving order")
	assert.Equal(t, savedOrder.Amount, int64(1), "wrong amount")
	assert.Equal(t, fetchedOrder.Amount, savedOrder.Amount, "wrong amount")
	assert.Equal(t, savedOrder.Status, domain.StatusDelivered, "wrong status")
	assert.Equal(t, fetchedOrder.Status, savedOrder.Status, "wrong amount")
}

func loadOrdersByBoard(ctx context.Context, t *testing.T, repository OrderRepositoryAdapter, board string, orders []string) {
	savedOrders, err := repository.LoadByBoard(ctx, board)
	var savedIds []string
	for _, order := range savedOrders {
		savedIds = append(savedIds, order.Id)
	}
	assert.Nil(t, err, "error when saving order")
	assert.ElementsMatch(t, savedIds, orders, "wrong orders load from bd compared to in memory")
}

func loadOrdersByPlayer(ctx context.Context, t *testing.T, repository OrderRepositoryAdapter, player string, orders []string) {
	savedOrders, err := repository.LoadByPlayer(ctx, player)
	var savedIds []string
	for _, order := range savedOrders {
		savedIds = append(savedIds, order.Id)
	}
	assert.Nil(t, err, "error when saving order")
	assert.ElementsMatch(t, savedIds, orders, "wrong response")
}
