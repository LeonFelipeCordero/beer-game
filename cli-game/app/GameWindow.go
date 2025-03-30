package app

import (
	"LeonFelipeCorder/beer-game/clitool/graph"
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

//type roleItem struct {
//	name        string
//	description string
//}
//
//func (i roleItem) Title() string       { return i.name }
//func (i roleItem) Description() string { return i.description }
//func (i roleItem) FilterValue() string { return i.name }

//func waitForBoardUpdate(updatesChannel chan *graph.Board) tea.Cmd {
//	return func() tea.Msg {
//		return <-updatesChannel
//	}
//}

type GameModel struct {
	player            graph.Player
	board             graph.Board
	boardGraphClient  graph.BoardGraphClient
	playerGraphClient graph.PlayerGraphClient
	//boardUpdates      chan *graph.Board
	playerUpdates chan *graph.Player
}

//type playerPhase int
//
//const (
//	roleSelection playerPhase = iota
//	waiting
//	roleSelected
//	onPlayerError
//)

func NewGameModel(boardId string, playerId string) *GameModel {
	ctx := context.Background()
	boardClient := graph.NewBoardGraphClient()
	playerClient := graph.NewPlayerGraphClient()

	board, err := boardClient.GetBoardByID(ctx, boardId)
	if err != nil {
		log.Println("fatal: ", err)
		os.Exit(1)
	}
	player, err := playerClient.GetPlayerByID(ctx, playerId)
	if err != nil {
		log.Println("fatal: ", err)
		os.Exit(1)
	}

	//boardUpdates, err := boardClient.SubscribeToBoard(ctx, boardId)
	//if err != nil {
	//	log.Fatalf("fatal, can not get board updates: %s\n", err.Error())
	//}
	playerUpdates, err := playerClient.SubscribeToPlayer(ctx, playerId)
	if err != nil {
		log.Fatalf("fatal, can not get player updates: %s\n", err.Error())
	}

	m := GameModel{
		board:             *board,
		player:            *player,
		boardGraphClient:  graph.NewBoardGraphClient(),
		playerGraphClient: graph.NewPlayerGraphClient(),
		//boardUpdates:      boardUpdates,
		playerUpdates: playerUpdates,
	}

	return &m
}

func (m *GameModel) Init() tea.Cmd {
	return tea.Batch(
		//waitForBoardUpdate(m.boardUpdates),
		waitForPlayerUpdate(m.playerUpdates),
	)
}

func (m *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case tea.KeyMsg:
	}

	return m, cmd
}

func (m *GameModel) View() string {
	var content string
	content = fmt.Sprintf("All set up %s-%s", m.board.Name, m.player.Role)

	return docStyle.Render(content)
}
