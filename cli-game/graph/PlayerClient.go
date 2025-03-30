package graph

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

type PlayerGraphClient struct {
}

type playerWrapper struct {
	Player    *Player `json:"player"`
	AddPlayer *Player `json:"addPlayer"`
	GetPlayer *Player `json:"getPlayer"`
}

type basicPlayerQuery struct {
	id          string
	role        string
	stock       float32
	backlog     float32
	weeklyOrder float32
	lastOrder   float64
	orders      []struct {
		id             string
		originalAmount float32
		state          string
		createdAt      time.Time
		receiver       struct {
			id string
		}
		sender struct {
			id string
		}
	}
}

var playerSubscribe struct {
	player basicPlayerQuery `graphql:"player(playerId: $id)"`
}

func NewPlayerGraphClient() PlayerGraphClient {
	return PlayerGraphClient{}
}

func (c PlayerGraphClient) AddPlayer(ctx context.Context, boardId, role string) (*Player, error) {
	var addPlayer struct {
		player basicPlayerQuery `graphql:"addPlayer(boardId: $boardId, role: $role)"`
	}
	variables := map[string]interface{}{
		"boardId": boardId,
		"role":    Role(role),
	}

	response, err := mutate(ctx, &addPlayer, variables)
	if err != nil {
		log.Printf("Error adding player to board: %v\n", err.Error())
		return nil, err
	}

	wrapper := playerWrapper{}
	err = json.Unmarshal(response, &wrapper)
	if err != nil {
		log.Printf("Error unparsing player response, %s: %v\n", wrapper, err.Error())
	}

	return wrapper.AddPlayer, nil
}

func (c PlayerGraphClient) GetPlayerByID(ctx context.Context, id string) (*Player, error) {
	var getPlayer struct {
		player basicPlayerQuery `graphql:"getPlayer(playerId: $playerId)"`
	}
	variables := map[string]interface{}{
		"playerId": id,
	}

	response, err := query(ctx, getPlayer, variables)
	if err != nil {
		log.Printf("Error getting player: %v\n", err.Error())
		return nil, err
	}

	wrapper := playerWrapper{}
	err = json.Unmarshal(response, &wrapper)
	if err != nil {
		log.Printf("Error unparsing player response, %s: %v\n", wrapper, err.Error())
	}

	return wrapper.GetPlayer, nil
}

func (c PlayerGraphClient) SubscribeToPlayer(ctx context.Context, playerId string) (chan *Player, error) {
	variables := map[string]interface{}{
		"id": playerId,
	}

	updates, err := subscribe(ctx, playerSubscribe, variables)
	if err != nil {
		log.Printf("Error subscribing to player: %v\n", err.Error())
	}

	playerUpdates := make(chan *Player)
	go func() {
		for message := range updates {
			wrapper := playerWrapper{}
			err = json.Unmarshal(*message, &wrapper)
			if err != nil {
				log.Printf("Error unmarshalling player: %v\n", err.Error())
			}
			playerUpdates <- wrapper.Player
		}
	}()

	return playerUpdates, nil
}
