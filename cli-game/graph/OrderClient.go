package graph

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

type OrderGraphClient struct {
}

type orderWrapper struct {
	Order       *Board `json:"order"`
	CreateBoard *Board `json:"createOrder"`
}

type basicOrderQuery struct {
	id             string
	originalAmount float32
	state          string
	_type          string `graphql:"type"`
	createdAt      time.Time
}

var newOrderSubscribe struct {
	board basicOrderQuery `graphql:"board(boardId: $id)"`
}

func NewOrderGraphClient() OrderGraphClient {
	return OrderGraphClient{}
}

func (c BoardGraphClient) CreateOrder(ctx context.Context, name string) (*Board, error) {
	var createBoard struct {
		board basicBoardQuery `graphql:"createBoard(name: $name)"`
	}
	variables := map[string]interface{}{
		"name": name,
	}

	response, err := mutate(ctx, &createBoard, variables)
	if err != nil {
		log.Printf("Error creating board: %v\n", err.Error())
		return nil, err
	}

	wrapper := boardWrapper{}
	err = json.Unmarshal(response, &wrapper)
	if err != nil {
		log.Printf("Error unparsing board response, %s: %v\n", wrapper, err.Error())
		return nil, err
	}

	return wrapper.CreateBoard, nil
}
