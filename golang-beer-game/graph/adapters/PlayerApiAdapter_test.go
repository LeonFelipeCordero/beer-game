package adapters

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
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
	playerApiAdapter := NewPlayerApiAdapter(playerService)

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
		assert.Equal(t, len(player.OrdersId), len(result.OrdersId), "wrong size")

		playerRepository.DeleteAll(ctx)
	})

	t.Run("a player should not be created if role taken", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)
		_, err := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		if err != nil {
			t.Error("There should not be errors")
		}

		player2, err := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		assert.Error(t, err, fmt.Sprintf("board name %s already exist", boardName))

		playerRepository.DeleteAll(ctx)
	})

	t.Run("a player should be part of a board", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardService.Create(ctx, boardName)
		player, err := playerApiAdapter.AddPlayer(ctx, board.Id, "RETAILER")
		if err != nil {
			t.Error("There should not be errors")
		}

		result, err := playerApiAdapter.Get(ctx, player.ID)

		assert.Equal(t, player.Name, result.Name, "wrong name")
		assert.Equal(t, player.Backlog, result.Backlog, "wrong backlog amount")
		assert.Equal(t, player.Stock, result.Stock, "wrong stock amount")
		assert.Equal(t, player.WeeklyOrder, result.WeeklyOrder, "wrong weekly order amount")
		assert.Equal(t, player.LastOrder, result.LastOrder, "wrong last order amount")
		assert.Equal(t, player.CPU, result.CPU, "wrong cpu")
		assert.Equal(t, len(player.OrdersId), len(result.OrdersId), "wrong size")

		playerRepository.DeleteAll(ctx)
	})
}
