package adapters

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	storage "github.com/LeonFelipeCordero/golang-beer-game/repositories/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BoardRepositoryAdapter struct {
	queries          *storage.Queries
	playerRepository PlayerRepositoryAdapter
}

func NewBoardRepository(queries *storage.Queries) BoardRepositoryAdapter {
	return BoardRepositoryAdapter{
		queries:          queries,
		playerRepository: NewPlayerRepository(queries),
	}
}

func (b *BoardRepositoryAdapter) Save(ctx context.Context, board domain.Board) (*domain.Board, error) {
	boardEntity, err := b.queries.SaveBoard(ctx, board.Name)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong creating new board %s", board.Name),
			err,
		)
	}

	return boardEntityToDomain(boardEntity), nil
}

func (b *BoardRepositoryAdapter) Get(ctx context.Context, boardId string) (*domain.Board, error) {
	boardEntity, err := b.queries.FindBoardById(ctx, pgtype.UUID{Bytes: uuid.MustParse(boardId), Valid: true})

	if err != nil && !isNotFound(err) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", boardId),
			err,
		)
	}

	board, err := b.complement(ctx, boardEntityToDomain(boardEntity))

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", boardId),
			err,
		)
	}

	return board, nil
}

func (b *BoardRepositoryAdapter) GetByName(ctx context.Context, name string) (*domain.Board, error) {
	boardEntity, err := b.queries.FindBoardByName(ctx, name)

	if err != nil && isNotFound(err) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", name),
			err,
		)
	}

	board, err := b.complement(ctx, boardEntityToDomain(boardEntity))

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", name),
			err,
		)
	}

	return board, nil
}

func (b *BoardRepositoryAdapter) Exist(ctx context.Context, name string) (bool, error) {
	boardEntity, err := b.GetByName(ctx, name)
	return boardEntity != nil, err
}

func (b *BoardRepositoryAdapter) DeleteAll(ctx context.Context) {
	err := b.queries.DeleteAllBoards(ctx)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("Something went wrong deleting all boards"),
			err,
		)
	}
}

func (b *BoardRepositoryAdapter) GetByPlayer(ctx context.Context, boardId string) (*domain.Board, error) {
	boardEntity, err := b.queries.FindBoardByPlayerId(ctx, pgtype.UUID{Bytes: uuid.MustParse(boardId), Valid: true})

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board by player %s", boardId),
			err,
		)
	}
	board, err := b.complement(ctx, boardEntityToDomain(boardEntity))

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting board %s", boardId),
			err,
		)
	}

	return board, nil
}

func (b *BoardRepositoryAdapter) GetActiveBoards(ctx context.Context) ([]domain.Board, error) {
	boardEntities, err := b.queries.GetRunningBoards(ctx)

	if err != nil && !isNotFound(err) {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting running boards"),
			err,
		)
	}

	var boards []domain.Board
	for _, boardEntity := range boardEntities {
		board := boardEntityToDomain(boardEntity)
		boards = append(boards, *board)
	}

	return boards, nil
}

func (b *BoardRepositoryAdapter) StartBoard(ctx context.Context, boardId string) error {
	err := b.queries.StartBoard(ctx, pgtype.UUID{Bytes: uuid.MustParse(boardId), Valid: true})
	if err != nil {
		return fmt.Errorf(
			fmt.Sprintf("Something went wrong starting board %s", boardId),
			err,
		)
	}
	return nil
}

func (b *BoardRepositoryAdapter) GetAvailableRoles(ctx context.Context, boardId string) ([]string, error) {
	roles, err := b.queries.GetAvailableRoles(ctx, pgtype.UUID{Bytes: uuid.MustParse(boardId), Valid: true})
	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong starting board %s", boardId),
			err,
		)
	}

	var result []string
	for _, role := range roles {
		result = append(result, role.String)
	}

	return result, nil
}

func (b *BoardRepositoryAdapter) complement(ctx context.Context, board *domain.Board) (*domain.Board, error) {
	player, err := b.playerRepository.GetPlayersByBoard(ctx, board.Id)

	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong fetching clients for board %s", board.Id),
			err,
		)
	}

	board.Players = player

	return board, nil
}

func boardEntityFromDomain(board domain.Board) storage.Board {
	timestamp := pgtype.Timestamp{}
	timestamp.Scan(board.CreatedAt)

	return storage.Board{
		BoardID:    pgtype.UUID{Bytes: uuid.MustParse(board.Id), Valid: true},
		Name:       board.Name,
		State:      string(board.State),
		IsFull:     board.Full,
		IsFinished: board.Finished,
		CreatedAt:  timestamp,
		UpdatedAt:  timestamp,
	}
}

func boardEntityToDomain(boardEntity storage.Board) *domain.Board {
	return &domain.Board{
		Id:        boardEntity.BoardID.String(),
		Name:      boardEntity.Name,
		State:     domain.State(boardEntity.State),
		Full:      boardEntity.IsFull,
		Finished:  boardEntity.IsFinished,
		CreatedAt: boardEntity.CreatedAt.Time,
	}
}
