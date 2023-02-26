package adapters

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j/entities"
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
	var boardNode = fromBoard(board)
	err := b.repository.Save(ctx, boardNode)
	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong creating new board %s", boardNode.Name),
			err,
		)
	}
	return toBoard(*boardNode), nil
}

func (b *BoardRepositoryAdapter) Get(ctx context.Context, id string) (*domain.Board, error) {
	query := "MATCH (b:BoardNode{uuid: $id}) return b"
	boardNode := &entities.BoardNode{}
	err := b.repository.Query(ctx, query, map[string]interface{}{"uuid": id}, boardNode)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", id),
			err,
		)
	}

	return toBoard(*boardNode), nil
}

func (b *BoardRepositoryAdapter) GetByName(ctx context.Context, name string) (*domain.Board, error) {
	query := "MATCH (b:BoardNode{name: $name}) return b"
	boardNode := &entities.BoardNode{}

	err := b.repository.Query(ctx, query, map[string]interface{}{"name": name}, boardNode)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", name),
			err,
		)
	}

	return toBoard(*boardNode), nil
}

func (b *BoardRepositoryAdapter) AddPlayer(ctx context.Context, boardId string, role string) (*domain.Player, error) {
	panic("not implemented")
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

func (b *BoardRepositoryAdapter) DeleteAll(ctx context.Context) {}

func (b *BoardRepositoryAdapter) GetByPlayer(ctx context.Context, id string) (*domain.Board, error) {
	//TODO implement me
	panic("implement me")
}

func fromBoard(board domain.Board) *entities.BoardNode {
	node := &entities.BoardNode{}
	node.Name = board.Name
	node.State = fmt.Sprint(board.State)
	node.Full = board.Full
	node.Finished = board.Finished
	node.CreatedAt = board.CreatedAt
	return node
}

func toBoard(node entities.BoardNode) *domain.Board {
	board := &domain.Board{}
	board.Id = node.UUID
	board.Name = node.Name
	board.State = toState(node.State)
	board.Full = node.Full
	board.Finished = node.Finished
	board.CreatedAt = node.CreatedAt
	return board
}

func toState(state string) domain.State {
	var result domain.State
	switch state {
	case string(domain.StateCreated):
		result = domain.StateCreated
	case string(domain.StateRunning):
		result = domain.StateRunning
	case string(domain.StateFinished):
		result = domain.StateFinished
	}
	return result
}
