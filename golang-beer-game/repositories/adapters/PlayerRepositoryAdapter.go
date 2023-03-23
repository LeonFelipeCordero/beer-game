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

type PlayerRepositoryAdapter struct {
	repository      neo4j.IRepository
	boardRepository ports.IBoardRepository
}

func NewPlayerRepository(
	repository neo4j.IRepository,
	boardRepository ports.IBoardRepository,
) ports.IPlayerRepository {
	return &PlayerRepositoryAdapter{
		repository:      repository,
		boardRepository: boardRepository,
	}
}

func (p PlayerRepositoryAdapter) AddPlayer(ctx context.Context, boardId string, player domain.Player) (*domain.Player, error) {
	board, err := p.boardRepository.Get(ctx, boardId)
	if err != nil {
		return nil, err
	}

	playerNode := &entities.PlayerNode{}
	playerNode.FromPlayer(player)
	boardNode := &entities.BoardNode{}
	boardNode.FromBoard(*board)
	playerNode.Board = boardNode

	p.repository.SaveDepth(ctx, playerNode)

	response := playerNode.ToPlayer()
	return &response, nil
}

func (p PlayerRepositoryAdapter) Get(ctx context.Context, id string) (*domain.Player, error) {
	entityId, _ := strconv.ParseInt(id, 0, 64)
	query := "MATCH (p:PlayerNode) WHERE ID(p) = $id RETURN p"
	playerNode := &entities.PlayerNode{}
	err := p.repository.Query(ctx, query, map[string]interface{}{"id": entityId}, playerNode)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", id),
			err,
		)
	}

	player := playerNode.ToPlayer()
	return &player, nil
}

func (p PlayerRepositoryAdapter) DeleteAll(ctx context.Context) {
	panic("implement me")
}

func (p PlayerRepositoryAdapter) Save(ctx context.Context, player domain.Player) (*domain.Player, error) {
	playerNode := &entities.PlayerNode{}
	playerNode.FromPlayer(player)
	err := p.repository.Save(ctx, playerNode)
	if err != nil {
		return nil, err
	}
	savePlayer := playerNode.ToPlayer()
	return &savePlayer, nil
}

func (p PlayerRepositoryAdapter) GetPlayersByBoard(ctx context.Context, boardId string) ([]domain.Player, error) {
	entityId, _ := strconv.ParseInt(boardId, 0, 64)
	query := "MATCH (p:PlayerNode)-[r:plays_in]->(b:BoardNode) WHERE ID(b) = $id RETURN p"

	playerNodes := &[]entities.PlayerNode{}

	err := p.repository.Query(ctx, query, map[string]interface{}{"id": entityId}, playerNodes)

	if err != nil && !isNotFound(err) {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", boardId),
			err,
		)
	}

	players := []domain.Player{}
	for _, playerNode := range *playerNodes {
		player := playerNode.ToPlayer()
		players = append(players, player)
	}

	return players, nil
}
