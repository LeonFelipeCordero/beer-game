package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/google/uuid"
)

type PlayerRepositoryAdapterFaker struct {
	players         map[string]domain.Player
	boardRepository ports.IBoardRepository
}

func NewPlayerRepositoryFaker(boardRepository ports.IBoardRepository) ports.IPlayerRepository {
	return &PlayerRepositoryAdapterFaker{
		players:         make(map[string]domain.Player),
		boardRepository: boardRepository,
	}
}

func (p *PlayerRepositoryAdapterFaker) AddPlayer(ctx context.Context, boardId string, player domain.Player) (*domain.Player, error) {
	id, _ := uuid.NewUUID()
	player.Id = id.String()
	p.players[id.String()] = player

	board, _ := p.boardRepository.Get(ctx, boardId)
	board.Players = append(board.Players, player)

	_, err := p.boardRepository.Save(ctx, *board)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (p *PlayerRepositoryAdapterFaker) Get(ctx context.Context, id string) (*domain.Player, error) {
	player := p.players[id]
	return &player, nil
}

func (p *PlayerRepositoryAdapterFaker) DeleteAll(ctx context.Context) {
	for key := range p.players {
		delete(p.players, key)
	}
}

func (p *PlayerRepositoryAdapterFaker) GetPlayersByBoard(ctx context.Context, boardId string) (*[]domain.Player, error) {
	board, _ := p.boardRepository.Get(ctx, boardId)
	return &board.Players, nil
}
