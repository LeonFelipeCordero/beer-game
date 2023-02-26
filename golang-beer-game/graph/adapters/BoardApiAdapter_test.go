package adapters

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	adapters2 "github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard(t *testing.T) {
	boardName := "test"

	boardRepository := adapters2.NewBoardRepositoryFaker()
	playerRepository := adapters2.NewPlayerRepositoryFaker(boardRepository)
	boardService := application.NewBoardService(boardRepository)
	playerService := application.NewPlayerService(playerRepository, boardService)
	boardApiAdapter := NewBoardApiAdapter(boardService)
	playerApiAdapter := NewPlayerApiAdapter(playerService, boardService)

	t.Run("a board should be created if name is not taken", func(t *testing.T) {
		ctx := context.Background()
		savedBoard, err := boardApiAdapter.Create(ctx, boardName)
		if err != nil {
			t.Error("There should not be errors")
		}

		result, err := boardApiAdapter.Get(ctx, savedBoard.ID)

		assert.Equal(t, savedBoard.Name, result.Name, "wrong name")
		assert.Equal(t, savedBoard.State, result.State, "wrong state")
		assert.Equal(t, savedBoard.Finished, result.Finished, "wrong finished")
		assert.Equal(t, savedBoard.Full, result.Full, "wrong full")
		assert.Equal(t, len(savedBoard.PlayersId), len(result.PlayersId), "wrong size")

		boardRepository.DeleteAll(ctx)
	})

	t.Run("a board should not be created if name is taken", func(t *testing.T) {
		ctx := context.Background()
		_, err := boardApiAdapter.Create(ctx, boardName)
		if err != nil {
			t.Error("There should not be errors")
		}

		_, err = boardApiAdapter.Create(ctx, boardName)
		assert.Error(t, err, fmt.Sprintf("board name %s already exist", boardName))

		secondBoard, err := boardApiAdapter.Create(ctx, boardName+"1")
		result, err := boardApiAdapter.Get(ctx, secondBoard.ID)

		assert.Equal(t, secondBoard.Name, result.Name, "wrong name")
		assert.Equal(t, secondBoard.State, result.State, "wrong state")
		assert.Equal(t, secondBoard.Finished, result.Finished, "wrong finished")
		assert.Equal(t, secondBoard.Full, result.Full, "wrong full")
		assert.Equal(t, len(secondBoard.PlayersId), len(result.PlayersId), "wrong size")

		boardRepository.DeleteAll(ctx)
	})

	t.Run("should return available roles", func(t *testing.T) {
		ctx := context.Background()
		board, _ := boardApiAdapter.Create(ctx, boardName)

		availableRoles, _ := boardApiAdapter.GetAvailableRoles(ctx, board.ID)
		assert.Equal(t, len(availableRoles), 3)
		_, _ = playerApiAdapter.AddPlayer(ctx, board.ID, "RETAILER")
		_, _ = playerApiAdapter.AddPlayer(ctx, board.ID, "WHOLESALER")
		_, _ = playerApiAdapter.AddPlayer(ctx, board.ID, "FACTORY")

		board, _ = boardApiAdapter.Get(ctx, board.ID)
		availableRoles, _ = boardApiAdapter.GetAvailableRoles(ctx, board.ID)
		assert.Equal(t, len(availableRoles), 0)

		boardRepository.DeleteAll(ctx)
		playerRepository.DeleteAll(ctx)
	})
}
