package adapters

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	adapters2 "github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer(t *testing.T) {
	boardName := "test"
	boardRepository := adapters2.NewBoardRepositoryFaker()
	playerRepository := adapters2.NewPlayerRepositoryFaker(boardRepository)
	boardService := application.NewBoardService(boardRepository)
	playerService := application.NewPlayerService(playerRepository, boardService)
	playerApiAdapter := NewPlayerApiAdapter(playerService, boardService)

	t.Run("a player should be created if role not taken", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)
		player, err := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		if err != nil {
			t.Error("There should not be errors", err)
		}

		result, err := playerApiAdapter.Get(ctx, player.ID)

		assert.Equal(t, player.Name, result.Name, "wrong name")
		assert.Equal(t, player.Backlog, result.Backlog, "wrong backlog amount")
		assert.Equal(t, player.Stock, result.Stock, "wrong stock amount")
		assert.Equal(t, player.WeeklyOrder, result.WeeklyOrder, "wrong weekly order amount")
		assert.Equal(t, player.LastOrder, result.LastOrder, "wrong last order amount")
		assert.Equal(t, player.CPU, result.CPU, "wrong cpu")
		assert.Equal(t, player.BoardId, board.Id, "wrong cpu")
		assert.Equal(t, len(player.OrdersId), len(result.OrdersId), "wrong size")

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
	})

	t.Run("a player should not be created if role taken", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)
		_, err := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		if err != nil {
			t.Error("There should not be errors")
		}

		_, err = playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		assert.Error(t, err, "Role RETAILER is already selected in the board")

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
	})

	t.Run("a player should be part of a board", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)
		player, err := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		if err != nil {
			t.Error("There should not be errors")
		}

		board, err = boardService.Get(ctx, player.BoardId)
		for _, boardPlayer := range board.Players {
			if boardPlayer.Id == player.ID {
				assert.Equal(t, player.Name, boardPlayer.Name, "wrong name")
				assert.Equal(t, player.Backlog, boardPlayer.Backlog, "wrong backlog amount")
				assert.Equal(t, player.Stock, boardPlayer.Stock, "wrong stock amount")
				assert.Equal(t, player.WeeklyOrder, boardPlayer.WeeklyOrder, "wrong weekly order amount")
				assert.Equal(t, player.LastOrder, boardPlayer.LastOrder, "wrong last order amount")
				assert.Equal(t, player.CPU, boardPlayer.Cpu, "wrong cpu")
				assert.Equal(t, len(player.OrdersId), len(boardPlayer.Orders), "wrong size")
			}
		}

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
	})

	t.Run("All players should be added to the board", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)
		player, _ := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		player, _ = playerApiAdapter.AddPlayer(ctx, board.Id, "FACTORY")
		player, _ = playerApiAdapter.AddPlayer(ctx, board.Id, "WHOLESALER")

		board, _ = boardService.Get(ctx, player.BoardId)
		assert.Equal(t, len(board.Players), 3, "players size is wrong")
		assert.Equal(t, board.State, domain.StateRunning, "players size is wrong")
		assert.Equal(t, board.Full, true, "players size is wrong")

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
	})

	t.Run("get all players by board board", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)

		retailer, _ := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		wholesaler, _ := playerApiAdapter.AddPlayer(ctx, board.Id, "WHOLESALER")
		factory, _ := playerApiAdapter.AddPlayer(ctx, board.Id, "FACTORY")

		players, err := playerApiAdapter.GetPlayersByBoard(ctx, board.Id)
		if err != nil {
			t.Error("There should not be errors")
		}
		assert.Equal(t, len(players), 3)
		for _, player := range players {
			if player.ID != retailer.ID || player.ID != wholesaler.ID || player.ID != factory.ID {
				assert.Error(t, errors.New(fmt.Sprintf("Player %s should not be part of board", player.ID)))
			}
		}

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
	})

	t.Run("get all players by board board", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)

		retailer, _ := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		assert.Equal(t, retailer.WeeklyOrder, 40)

		response, err := playerApiAdapter.UpdateWeeklyOrder(ctx, retailer.ID, 10000)
		if err != nil {
			t.Error("There should not be errors")
		}

		assert.Equal(t, *response.Message, "weekly order updated")
		assert.Equal(t, *response.Status, 200)

		player, _ := playerApiAdapter.Get(ctx, retailer.ID)
		assert.Equal(t, player.WeeklyOrder, 10000)

		playerRepository.DeleteAll(ctx)
		boardRepository.DeleteAll(ctx)
	})
}
