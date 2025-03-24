package adapters

import (
	"context"
	"fmt"

	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	storage "github.com/LeonFelipeCordero/golang-beer-game/repositories/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PlayerRepositoryAdapter struct {
	queries         *storage.Queries
	orderRepository OrderRepositoryAdapter
}

func NewPlayerRepository(queries *storage.Queries) PlayerRepositoryAdapter {
	return PlayerRepositoryAdapter{
		queries:         queries,
		orderRepository: NewOrderRepository(queries),
	}
}

func (p PlayerRepositoryAdapter) AddPlayer(ctx context.Context, boardId string, player domain.Player) (*domain.Player, error) {
	params := storage.SavePlayerParams{
		Name:        player.Name,
		Role:        string(player.Role),
		Stock:       player.Stock,
		Backlog:     player.Backlog,
		WeeklyOrder: player.WeeklyOrder,
		LastOrder:   player.LastOrder,
		Cpu:         player.Cpu,
		BoardID:     pgtype.UUID{Bytes: uuid.MustParse(boardId), Valid: true},
	}
	playerEntity, err := p.queries.SavePlayer(ctx, params)

	if err != nil {
		return nil, err
	}

	return playerEntityToDomain(playerEntity), nil
}

func (p PlayerRepositoryAdapter) Get(ctx context.Context, id string) (*domain.Player, error) {
	playerId := pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}
	playerEntity, err := p.queries.FindPlayerById(ctx, playerId)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		} else {
			return nil, fmt.Errorf(
				fmt.Sprintf("Something went wrong getting player %s", id),
				err,
			)
		}
	}

	player := playerEntityToDomain(playerEntity)
	orders, err := p.orderRepository.LoadByPlayer(ctx, id)
	player.Orders = orders

	return player, nil
}

func (p PlayerRepositoryAdapter) DeleteAll(ctx context.Context) {
	panic("implement me")
}

func (p PlayerRepositoryAdapter) Save(ctx context.Context, player domain.Player) (*domain.Player, error) {
	params := storage.UpdatePlayerNumbersParams{
		PlayerID:    pgtype.UUID{Bytes: uuid.MustParse(player.Id), Valid: true},
		Stock:       player.Stock,
		Backlog:     player.Backlog,
		WeeklyOrder: player.WeeklyOrder,
		LastOrder:   player.LastOrder,
	}
	err := p.queries.UpdatePlayerNumbers(ctx, params)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (p PlayerRepositoryAdapter) GetPlayersByBoard(ctx context.Context, boardId string) ([]domain.Player, error) {
	playerEntities, err := p.queries.FindPlayerByBoardId(ctx, pgtype.UUID{Bytes: uuid.MustParse(boardId), Valid: true})
	if err != nil && !isNotFound(err) {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting player by board %s", boardId),
			err,
		)
	}

	var players []domain.Player
	for _, playerEntity := range playerEntities {
		player := playerEntityToDomain(playerEntity)
		players = append(players, *player)
	}

	return players, nil
}

func (p PlayerRepositoryAdapter) GetByRoleAndBoardId(ctx context.Context, boardId, role string) (*domain.Player, error) {
	id := pgtype.UUID{}
	id.Scan(boardId)
	params := storage.GetPlayerByRoleAndBoardIdParams{
		BoardID: id,
		Role:    role,
	}
	playerEntity, err := p.queries.GetPlayerByRoleAndBoardId(ctx, params)

	if err != nil {
		return nil, err
	}

	return playerEntityToDomain(playerEntity), nil
}

func playerEntityFromDomain(player domain.Player) storage.Player {
	id := pgtype.UUID{}
	id.Scan(player.Id)
	timestamp := pgtype.Timestamp{}
	timestamp.Scan(player.CreatedAt)

	return storage.Player{
		PlayerID:    id,
		Name:        player.Name,
		Role:        string(player.Role),
		Stock:       player.Stock,
		Backlog:     player.Backlog,
		WeeklyOrder: player.WeeklyOrder,
		LastOrder:   player.LastOrder,
		CreatedAt:   timestamp,
		UpdatedAt:   timestamp,
	}
}

func playerEntityToDomain(playerEntity storage.Player) *domain.Player {
	return &domain.Player{
		Id:          playerEntity.PlayerID.String(),
		Name:        playerEntity.Name,
		Role:        domain.Role(playerEntity.Role),
		Stock:       playerEntity.Stock,
		Backlog:     playerEntity.Backlog,
		WeeklyOrder: playerEntity.WeeklyOrder,
		LastOrder:   playerEntity.LastOrder,
		CreatedAt:   playerEntity.CreatedAt.Time,
	}
}
