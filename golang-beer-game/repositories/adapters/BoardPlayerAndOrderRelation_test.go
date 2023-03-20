package adapters

import (
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
	"testing"
)

func TestBoardPlayerAndOrder(t *testing.T) {
	neo4j.ConfigureDatabase()
	//ctx := context.Background()
	//repository := neo4j.NewRepository()
	//boardRepository := NewBoardRepository(repository)
	//playerRepository := NewPlayerRepository(repository, boardRepository)
	//orderRepository := NewOrderRepository(repository, playerRepository)
	t.Run("create all relations between boards and player easily", func(t *testing.T) {
		//board := createBoard(ctx, boardRepository)
		//retailer := createPlayer(ctx, playerRepository, board, "RETAILER")
		//wholesaler := createPlayer(ctx, playerRepository, board, "WHOLESALER")

	})
}
