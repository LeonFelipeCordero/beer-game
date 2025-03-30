package graph

import (
	"context"
	"encoding/json"
	"log"
)

type BoardGraphClient struct {
}

type boardWrapper struct {
	Board          *Board `json:"board"`
	CreateBoard    *Board `json:"createBoard"`
	GetBoardByName *Board `json:"getBoardByName"`
	GetBoard       *Board `json:"getBoard"`
}

type basicBoardQuery struct {
	id      string
	name    string
	state   string
	full    bool
	players []struct {
		id   string
		role string
	}
	availableRoles []string
}

var boardSubscribe struct {
	board basicBoardQuery `graphql:"board(boardId: $id)"`
}

func NewBoardGraphClient() BoardGraphClient {
	return BoardGraphClient{}
}

func (c BoardGraphClient) CreateBoard(ctx context.Context, name string) (*Board, error) {
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

func (c BoardGraphClient) GetBoardByName(ctx context.Context, name string) (*Board, error) {
	var getBoardByName struct {
		board basicBoardQuery `graphql:"getBoardByName(name: $name)"`
	}
	variables := map[string]interface{}{
		"name": name,
	}

	response, err := query(ctx, getBoardByName, variables)
	if err != nil {
		log.Printf("Error getting board: %v\n", err.Error())
		return nil, err
	}

	wrapper := boardWrapper{}
	err = json.Unmarshal(response, &wrapper)
	if err != nil {
		log.Printf("Error unparsing board response, %s: %v\n", wrapper, err.Error())
	}

	return wrapper.GetBoardByName, nil
}

func (c BoardGraphClient) GetBoardByID(ctx context.Context, id string) (*Board, error) {
	var getBoard struct {
		board basicBoardQuery `graphql:"getBoard(id: $id)"`
	}
	variables := map[string]interface{}{
		"id": id,
	}

	response, err := query(ctx, getBoard, variables)
	if err != nil {
		log.Printf("Error getting board: %v\n", err.Error())
		return nil, err
	}

	wrapper := boardWrapper{}
	err = json.Unmarshal(response, &wrapper)
	if err != nil {
		log.Printf("Error unparsing board response, %s: %v\n", wrapper, err.Error())
	}

	return wrapper.GetBoard, nil
}

func (c BoardGraphClient) SubscribeToBoard(ctx context.Context, boardID string) (chan *Board, error) {
	variables := map[string]interface{}{
		"id": boardID,
	}

	updates, err := subscribe(ctx, boardSubscribe, variables)
	if err != nil {
		log.Printf("Error subscribing to board: %v\n", err.Error())
	}

	boardUpdates := make(chan *Board)
	go func() {
		for message := range updates {
			wrapper := boardWrapper{}
			err = json.Unmarshal(*message, &wrapper)
			if err != nil {
				log.Printf("Error unmarshalling board: %v\n", err.Error())
			}
			boardUpdates <- wrapper.Board
		}
	}()

	return boardUpdates, nil
}
