package app

import (
	"LeonFelipeCorder/beer-game/clitool/graph"
	tea "github.com/charmbracelet/bubbletea"
)

func waitForBoardUpdate(updatesChannel chan *graph.Board) tea.Cmd {
	return func() tea.Msg {
		return <-updatesChannel
	}
}

func waitForPlayerUpdate(updatesChannel chan *graph.Player) tea.Cmd {
	return func() tea.Msg {
		return <-updatesChannel
	}
}
