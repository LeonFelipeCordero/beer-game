package adapters

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	storage "github.com/LeonFelipeCordero/golang-beer-game/repositories/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderRepositoryAdapter struct {
	queries *storage.Queries
}

func NewOrderRepository(queries *storage.Queries) OrderRepositoryAdapter {
	return OrderRepositoryAdapter{
		queries: queries,
	}
}

func (o OrderRepositoryAdapter) Save(ctx context.Context, order domain.Order) (*domain.Order, error) {
	if order.OrderType == domain.OrderTypePlayerOrder {
		params := storage.SaveOrderParams{
			Amount:     order.Amount,
			Type:       string(order.OrderType),
			SenderID:   pgtype.UUID{Bytes: uuid.MustParse(order.Sender), Valid: true},
			ReceiverID: pgtype.UUID{Bytes: uuid.MustParse(order.Receiver), Valid: true},
		}
		orderEntity, err := o.queries.SaveOrder(ctx, params)
		if err != nil {
			return nil, err
		}
		return orderEntityToDomain(orderEntity), nil
	} else {
		params := storage.SaveCpuOrderParams{
			Amount:   order.Amount,
			Type:     string(order.OrderType),
			SenderID: pgtype.UUID{Bytes: uuid.MustParse(order.Sender), Valid: true},
		}
		orderEntity, err := o.queries.SaveCpuOrder(ctx, params)
		if err != nil {
			return nil, err
		}
		return orderEntityToDomain(orderEntity), nil
	}
}

func (o OrderRepositoryAdapter) Get(ctx context.Context, orderId string) (*domain.Order, error) {
	pgId := pgtype.UUID{Bytes: uuid.MustParse(orderId), Valid: true}
	orderEntity, err := o.queries.FindOrderById(ctx, pgId)
	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting order %s", orderId),
			err,
		)
	}
	return orderEntityToDomain(orderEntity), nil
}

func (o OrderRepositoryAdapter) DeleteAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepositoryAdapter) LoadByBoard(ctx context.Context, boardId string) ([]domain.Order, error) {
	pgId := pgtype.UUID{Bytes: uuid.MustParse(boardId), Valid: true}
	orderEntities, err := o.queries.FindOrderByBoardId(ctx, pgId)
	if err != nil && !isNotFound(err) {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting order by board %s", boardId),
			err,
		)
	}
	var orders []domain.Order
	for _, orderEntity := range orderEntities {
		order := orderEntityToDomain(orderEntity)
		orders = append(orders, *order)
	}
	return orders, nil
}

func (o OrderRepositoryAdapter) LoadByPlayer(ctx context.Context, playerId string) ([]domain.Order, error) {
	pgId := pgtype.UUID{Bytes: uuid.MustParse(playerId), Valid: true}
	orderEntities, err := o.queries.FindOrderByPlayerId(ctx, pgId)
	if err != nil && !isNotFound(err) {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong getting order by player %s", playerId),
			err,
		)
	}
	var orders []domain.Order
	for _, orderEntity := range orderEntities {
		order := orderEntityToDomain(orderEntity)
		orders = append(orders, *order)
	}
	return orders, nil
}

func (o OrderRepositoryAdapter) MarkAsFilled(ctx context.Context, orderId string, amount int64) (*domain.Order, error) {
	params := storage.MarkAsFilledParams{
		Amount:  amount,
		OrderID: pgtype.UUID{Bytes: uuid.MustParse(orderId), Valid: true},
	}
	orderEntity, err := o.queries.MarkAsFilled(ctx, params)
	if err != nil && !isNotFound(err) {
		return nil, fmt.Errorf(
			fmt.Sprintf("Something went wrong market order %s as filled", orderId),
			err,
		)
	}
	return orderEntityToDomain(orderEntity), nil
}

func orderEntityFromDomain(order domain.Order) storage.Order {
	timestamp := pgtype.Timestamp{}
	timestamp.Scan(order.CreatedAt)

	return storage.Order{
		OrderID:        pgtype.UUID{Bytes: uuid.MustParse(order.Id), Valid: true},
		Amount:         order.Amount,
		OriginalAmount: order.OriginalAmount,
		Type:           string(order.OrderType),
		State:          string(order.Status),
		SenderID:       pgtype.UUID{Bytes: uuid.MustParse(order.Sender), Valid: true},
		ReceiverID:     pgtype.UUID{Bytes: uuid.MustParse(order.Receiver), Valid: true},
		CreatedAt:      timestamp,
		UpdatedAt:      timestamp,
	}
}

func orderEntityToDomain(orderEntity storage.Order) *domain.Order {
	return &domain.Order{
		Id:             orderEntity.OrderID.String(),
		Amount:         orderEntity.Amount,
		OriginalAmount: orderEntity.OriginalAmount,
		OrderType:      domain.OrderType(orderEntity.Type),
		Status:         domain.Status(orderEntity.State),
		Sender:         orderEntity.SenderID.String(),
		Receiver:       orderEntity.ReceiverID.String(),
		CreatedAt:      orderEntity.CreatedAt.Time,
	}
}
