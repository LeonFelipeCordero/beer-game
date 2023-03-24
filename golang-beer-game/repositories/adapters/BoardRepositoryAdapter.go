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

type BoardRepositoryAdapter struct {
	repository neo4j.IRepository
}

func NewBoardRepository(repository neo4j.IRepository) ports.IBoardRepository {
	return &BoardRepositoryAdapter{
		repository: repository,
	}
}

func (b *BoardRepositoryAdapter) Save(ctx context.Context, board domain.Board) (*domain.Board, error) {
	boardNode := &entities.BoardNode{}
	boardNode.FromBoard(board)
	err := b.repository.SaveDepth(ctx, boardNode)
	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong creating new board %s", boardNode.Name),
			err,
		)
	}
	return boardNode.ToBoard(), nil
}

func (b *BoardRepositoryAdapter) Get(ctx context.Context, id string) (*domain.Board, error) {
	entityId, _ := strconv.ParseInt(id, 0, 64)
	boardNode := &entities.BoardNode{}
	err := b.repository.LoadDepth(ctx, entityId, boardNode)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", id),
			err,
		)
	}

	return boardNode.ToBoard(), nil
}

func (b *BoardRepositoryAdapter) GetByName(ctx context.Context, name string) (*domain.Board, error) {
	query := "MATCH (b:BoardNode{name: $name}) RETURN b"
	boardNode := &entities.BoardNode{}

	err := b.repository.Query(ctx, query, map[string]interface{}{"name": name}, boardNode)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", name),
			err,
		)
	}

	return boardNode.ToBoard(), nil
}

func (b *BoardRepositoryAdapter) Exist(ctx context.Context, name string) (bool, error) {
	query := "MATCH (b:BoardNode{name: $name}) RETURN count(b) as count"

	result, err := b.repository.QueryRaw(ctx, query, map[string]interface{}{
		"name": name,
	})

	if err != nil {
		return true, fmt.Errorf(
			fmt.Sprintf("Something went wrong validating board %s", name),
			err,
		)
	}

	return result[0][0].(int64) != 0, nil
}

func (b *BoardRepositoryAdapter) DeleteAll(ctx context.Context) {
	query := "MATCH (n) detach delete n"

	_, err := b.repository.QueryRaw(ctx, query, map[string]interface{}{})

	if err != nil {
		fmt.Println(
			fmt.Sprintf("Something went wrong deleting all boards"),
			err,
		)
	}

}

func (b *BoardRepositoryAdapter) GetByPlayer(ctx context.Context, id string) (*domain.Board, error) {
	entityId, _ := strconv.ParseInt(id, 0, 64)
	query := `
		MATCH (p:PlayerNode)-[r:plays_in]->(b:BoardNode) WHERE ID(p) = $id
		CALL {
			WITH b
			MATCH (p2:PlayerNode)-[r2:plays_in]->(b2:BoardNode) WHERE b2.name = b.name RETURN b2,r2,p2
		}
		RETURN b2,r2,p2
	`

	boardNode := &entities.BoardNode{}

	err := b.repository.Query(ctx, query, map[string]interface{}{"id": entityId}, boardNode)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board by player %s", id),
			err,
		)
	}

	return boardNode.ToBoard(), nil
}
