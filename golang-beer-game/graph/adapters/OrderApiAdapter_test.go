package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	adapters2 "github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder(t *testing.T) {
	boardName := "test"
	boardRepository := adapters2.NewBoardRepositoryFaker()
	playerRepository := adapters2.NewPlayerRepositoryFaker(boardRepository)
	orderRepository := adapters2.NewOrderRepositoryFaker()
	boardService := application.NewBoardService(boardRepository)
	playerService := application.NewPlayerService(playerRepository, boardService)
	orderService := application.NewOrderService(orderRepository, boardService, playerService)
	orderApiAdapter := NewOrderApiAdapter(orderService, boardService, playerService)

	t.Run("order should created and contain the correct data", func(t *testing.T) {
		ctx := context.Background()
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

		assert.Equal(t, retailerOrder.Amount, 40, "wrong amount for retailer order")
		assert.Equal(t, wholesalerOrder.Amount, 600, "wrong amount for wholesaler order")
		assert.Equal(t, retailerOrder.OriginalAmount, 40, "wrong original amount for retailer order")
		assert.Equal(t, wholesalerOrder.OriginalAmount, 600, "wrong original amount for wholesaler amount")

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
		orderRepository.DeleteAll(ctx)
	})

	t.Run("players should contains the correct orders", func(t *testing.T) {
		ctx := context.Background()
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

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
		orderRepository.DeleteAll(ctx)
	})

	t.Run("orders should be delivered", func(t *testing.T) {
		ctx := context.Background()
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

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
		orderRepository.DeleteAll(ctx)
	})
}
