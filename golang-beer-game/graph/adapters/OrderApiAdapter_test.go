package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	adapters2 "github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	testingutil "github.com/LeonFelipeCordero/golang-beer-game/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder(t *testing.T) {
	ctx := context.Background()
	boardName := "test"

	queries := testingutil.SetupDatabaseConnection(ctx)

	streamers, eventChan := events.CreateEventBus()
	go events.EventHandler(streamers, eventChan)

	boardRepository := adapters2.NewBoardRepository(queries)
	playerRepository := adapters2.NewPlayerRepository(queries)
	orderRepository := adapters2.NewOrderRepository(queries)
	boardService := application.NewBoardService(boardRepository, eventChan)
	playerService := application.NewPlayerService(playerRepository, boardService, eventChan)
	orderService := application.NewOrderService(orderRepository, playerService, boardService, eventChan)
	orderApiAdapter := NewOrderApiAdapter(orderService, boardService)

	t.Run("order should created and contain the correct data", func(t *testing.T) {
		testingutil.Clean(ctx, queries)
		board, _ := boardService.Create(ctx, boardName)
		retailer, _ := playerService.AddPlayer(ctx, board.Id, "RETAILER")
		wholesaler, _ := playerService.AddPlayer(ctx, board.Id, "WHOLESALER")
		playerService.AddPlayer(ctx, board.Id, "FACTORY")

		retailerOrder, err := orderApiAdapter.CreateOrder(ctx, retailer.Id)
		if err != nil {
			t.Error("There should not be errors")
		}
		wholesalerOrder, err := orderApiAdapter.CreateOrder(ctx, wholesaler.Id)
		if err != nil {
			t.Error("There should not be errors")
		}

		savedRetailerOrder, _ := orderService.Get(ctx, retailerOrder.ID)
		savedWholesalerOrder, _ := orderService.Get(ctx, wholesalerOrder.ID)

		assert.Equal(t, retailerOrder.ID, savedRetailerOrder.Id, "wrong retailer order ids")
		assert.Equal(t, wholesalerOrder.ID, savedWholesalerOrder.Id, "wrong wholesalers orders ids")

		assert.Equal(t, wholesalerOrder.Amount, savedWholesalerOrder.Amount, "wrong amount for wholesalers order")
		assert.Equal(t, retailerOrder.OriginalAmount, savedRetailerOrder.OriginalAmount, "wrong original amount for retailers order")
		assert.Equal(t, wholesalerOrder.OriginalAmount, savedWholesalerOrder.OriginalAmount, "wrong original amount for wholesalers amount")

		assert.Equal(t, retailerOrder.Amount, int64(40), "wrong amount for retailer order")
		assert.Equal(t, wholesalerOrder.Amount, int64(600), "wrong amount for wholesaler order")
		assert.Equal(t, retailerOrder.OriginalAmount, int64(40), "wrong original amount for retailer order")
		assert.Equal(t, wholesalerOrder.OriginalAmount, int64(600), "wrong original amount for wholesaler amount")
	})

	t.Run("players should contains the correct orders", func(t *testing.T) {
		testingutil.Clean(ctx, queries)
		board, _ := boardService.Create(ctx, boardName)
		retailer, _ := playerService.AddPlayer(ctx, board.Id, "RETAILER")
		wholesaler, _ := playerService.AddPlayer(ctx, board.Id, "WHOLESALER")
		playerService.AddPlayer(ctx, board.Id, "FACTORY")

		retailerOrder, _ := orderApiAdapter.CreateOrder(ctx, retailer.Id)
		wholesalerOrder, _ := orderApiAdapter.CreateOrder(ctx, wholesaler.Id)

		savedRetailer, _ := playerService.Get(ctx, retailerOrder.ReceiverId)
		savedWholesaler, _ := playerService.Get(ctx, wholesalerOrder.ReceiverId)

		assert.Equal(t, len(savedRetailer.Orders), 1, "wrong retailer orders")
		assert.Equal(t, len(savedWholesaler.Orders), 2, "wrong wholesaler orders")

		assert.Equal(t, savedRetailer.Orders[0].Id, retailerOrder.ID, "wrong retailer orders")
		assert.Equal(t, savedWholesaler.Orders[0].Id, retailerOrder.ID, "wrong wholesaler orders")
		assert.Equal(t, savedWholesaler.Orders[1].Id, wholesalerOrder.ID, "wrong wholesaler orders")
	})

	t.Run("orders should be delivered", func(t *testing.T) {
		testingutil.Clean(ctx, queries)
		board, _ := boardService.Create(ctx, boardName)
		retailer, _ := playerService.AddPlayer(ctx, board.Id, "RETAILER")
		wholesaler, _ := playerService.AddPlayer(ctx, board.Id, "WHOLESALER")
		playerService.AddPlayer(ctx, board.Id, "FACTORY")

		retailerOrder, _ := orderApiAdapter.CreateOrder(ctx, retailer.Id)

		response, err := orderApiAdapter.DeliverOrder(ctx, retailerOrder.ID, retailerOrder.Amount-1)
		if err != nil {
			t.Error("there should not be errors")
		}

		savedRetailer, _ := playerService.Get(ctx, retailerOrder.ReceiverId)
		savedWholesaler, _ := playerService.Get(ctx, retailerOrder.SenderId)
		savedOrder, _ := orderService.Get(ctx, retailerOrder.ID)

		assert.Equal(t, *response.Status, 200, "wrong response status")
		assert.Equal(t, *response.Message, "order delivered", "wrong response message")

		assert.Equal(t, savedOrder.Amount, retailerOrder.Amount-1, "wrong order amount")
		assert.Equal(t, savedOrder.Status, domain.StatusDelivered, "wrong order state")

		assert.Equal(t, savedRetailer.Stock, retailer.Stock+savedOrder.Amount, "wrong retailer stock")
		assert.Equal(t, savedWholesaler.Stock, wholesaler.Stock-savedOrder.Amount, "wrong wholesaler stock")
	})
}
