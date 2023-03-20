package adapters

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBoardAndPlayer(t *testing.T) {
	neo4j.ConfigureDatabase()
	ctx := context.Background()
	repository := neo4j.NewRepository()
	boardRepository := NewBoardRepository(repository)
	playerRepository := NewPlayerRepository(repository, boardRepository)
	t.Run("create all relations between boards and player easily", func(t *testing.T) {
		board := createBoard(ctx, boardRepository)
		retailer := createPlayer(ctx, playerRepository, board, "RETAILER")
		wholesaler := createPlayer(ctx, playerRepository, board, "WHOLESALER")

		checkIfBoardExit(ctx, t, boardRepository, board)
		checkBoard(ctx, t, boardRepository, board)
		checkBoardByPlayer(ctx, t, boardRepository, board, retailer)
		//checkBoardByName(ctx, t, boardRepository, board)
		checkPlayerByBoard(ctx, t, playerRepository, retailer, board)
		checkPlayerByBoard(ctx, t, playerRepository, wholesaler, board)
		checkPlayer(ctx, t, playerRepository, retailer)
		updatePlayerAndValidate(ctx, t, playerRepository, retailer)
	})
}

func createBoard(ctx context.Context, repository ports.IBoardRepository) *domain.Board {
	board := &domain.Board{
		Name:      "test",
		State:     "CREATED",
		Full:      false,
		Finished:  false,
		CreatedAt: time.Now().UTC(),
	}

	var savedBoard, err = repository.Save(ctx, *board)

	if err != nil {
		fmt.Printf("%e", err)
		panic("Should not fail")
	}
	return savedBoard
}

func getBoard(ctx context.Context, repository ports.IBoardRepository, board *domain.Board) *domain.Board {
	board, err := repository.Get(ctx, board.Id)
	if err != nil {
		fmt.Printf("%e", err)
		panic("Should not fail")
	}
	return board
}

func checkBoard(ctx context.Context, t *testing.T, repository ports.IBoardRepository, board *domain.Board) {
	savedBoard := getBoard(ctx, repository, board)
	validateBoard(t, *board, *savedBoard)
}

func checkBoardByPlayer(ctx context.Context, t *testing.T, repository ports.IBoardRepository, board *domain.Board, retailer *domain.Player) {
	savedBoard, err := repository.GetByPlayer(ctx, retailer.Id)

	if err != nil {
		fmt.Printf("%e", err)
	}

	validateBoard(t, *board, *savedBoard)
}

func checkIfBoardExit(ctx context.Context, t *testing.T, repository ports.IBoardRepository, board *domain.Board) {
	exist, err := repository.Exist(ctx, board.Name)

	if err != nil {
		fmt.Printf("%e", err)
	}

	assert.Equal(t, exist, true, "wrong answer")
}

//func checkBoardByName(ctx context.Context, t *testing.T, repository ports.IBoardRepository, board *domain.Board) {
//	savedBoard, err := repository.GetByName(ctx, board.Name)
//
//	if err != nil {
//		fmt.Printf("%e", err)
//	}
//
//	validateBoard(t, *board, *savedBoard)
//}

func createPlayer(ctx context.Context, repository ports.IPlayerRepository, board *domain.Board, role domain.Role) *domain.Player {
	player := domain.Player{
		Name:        string(role),
		Role:        role,
		Stock:       40,
		Backlog:     40,
		WeeklyOrder: 40,
		LastOrder:   40,
		Cpu:         false,
	}

	savedPlayer, err := repository.AddPlayer(ctx, board.Id, player)

	if err != nil {
		fmt.Printf("%e", err)
		panic("Should not fail")
	}
	return savedPlayer
}

func checkPlayerByBoard(ctx context.Context, t *testing.T, repository ports.IPlayerRepository, target *domain.Player, board *domain.Board) {
	players, _ := repository.GetPlayersByBoard(ctx, board.Id)
	var checked = false
	for _, player := range players {
		if player.Role == target.Role {
			checked = true
			validatePlayers(t, *target, player)
		}
	}
	assert.Equal(t, checked, true, "there was no player to compare to")
}

func updatePlayerAndValidate(ctx context.Context, t *testing.T, repository ports.IPlayerRepository, player *domain.Player) {
	player.Stock = 100
	player.Backlog = 101
	player.WeeklyOrder = 102
	player.LastOrder = 103

	savedPlayer, err := repository.Save(ctx, *player)
	assert.Nil(t, err, "there should no be errors")
	assert.Equal(t, savedPlayer.Stock, 100, "stock is wrong")
	assert.Equal(t, savedPlayer.Backlog, 101, "backlog is wrong")
	assert.Equal(t, savedPlayer.WeeklyOrder, 102, "weekly order is wrong")
	assert.Equal(t, savedPlayer.LastOrder, 103, "last order is wrong")
}

func checkPlayer(ctx context.Context, t *testing.T, repository ports.IPlayerRepository, player *domain.Player) {
	savedPlayer, _ := repository.Get(ctx, player.Id)
	validatePlayers(t, *player, *savedPlayer)
}

func validateBoard(t *testing.T, board1 domain.Board, board2 domain.Board) {
	assert.Equal(t, board1.Id, board2.Id)
	assert.Equal(t, board1.Name, board2.Name)
	assert.Equal(t, board1.Full, board2.Full)
	assert.Equal(t, board1.Finished, board2.Finished)
	assert.Equal(t, board1.State, board2.State)
}

func validatePlayers(t *testing.T, player1 domain.Player, player2 domain.Player) {
	assert.Equal(t, player1.Id, player2.Id)
	assert.Equal(t, player1.Name, player2.Name)
	assert.Equal(t, player1.Stock, player2.Stock)
	assert.Equal(t, player1.Backlog, player2.Backlog)
	assert.Equal(t, player1.WeeklyOrder, player2.WeeklyOrder)
	assert.Equal(t, player1.LastOrder, player2.LastOrder)
	assert.Equal(t, player1.Cpu, player2.Cpu)
}
