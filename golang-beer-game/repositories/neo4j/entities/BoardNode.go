package entities

import (
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/mindstand/gogm/v2"
	"strconv"
	"time"
)

type BoardNode struct {
	gogm.BaseNode

	Name      string        `gogm:"name=name;unique"`
	State     string        `gogm:"name=state"`
	Full      bool          `gogm:"name=full"`
	Finished  bool          `gogm:"name=finished"`
	CreatedAt time.Time     `gogm:"name=created_at"`
	Players   []*PlayerNode `gogm:"direction=incoming;relationship=plays_in"`
}

func (b *BoardNode) FromBoard(board domain.Board) {
	if board.Id != "" {
		id, _ := strconv.ParseInt(board.Id, 0, 64)
		b.Id = &id
	}
	b.Name = board.Name
	b.State = string(board.State)
	b.Full = board.Full
	b.Finished = board.Finished
	b.CreatedAt = board.CreatedAt
	b.Players = mapToPlayersNode(board.Players)
}

func (b *BoardNode) ToBoard() *domain.Board {
	board := &domain.Board{}
	board.Id = strconv.FormatInt(*b.Id, 10)
	board.Name = b.Name
	board.State = toState(b.State)
	board.Full = b.Full
	board.Finished = b.Finished
	board.CreatedAt = b.CreatedAt
	board.Players = mapFromPlayersNode(b.Players)
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

func mapFromPlayersNode(playersNode []*PlayerNode) []domain.Player {
	players := []domain.Player{}
	for _, playerNode := range playersNode {
		player := playerNode.ToPlayer()
		players = append(players, player)
	}
	return players
}

func mapToPlayersNode(players []domain.Player) []*PlayerNode {
	nodePlayers := []*PlayerNode{}
	for _, player := range players {
		nodePlayer := PlayerNode{}
		nodePlayer.FromPlayer(player)
		nodePlayers = append(nodePlayers, &nodePlayer)
	}
	return nodePlayers
}
