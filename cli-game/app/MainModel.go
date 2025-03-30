package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type window int

const (
	boardWindow window = iota
	playerWindow
	gameWindow
)

type WindowModel struct {
	window      window
	boardModel  tea.Model
	playerModel tea.Model
	gameModel   tea.Model
}

func NewMainModel() WindowModel {
	//return WindowModel{window: boardWindow, boardModel: NewBoardModel()}
	return WindowModel{window: gameWindow, gameModel: NewGameModel("025314ac-cc10-4d90-a098-d97468cccda5", os.Getenv("ID"))}
}

func (m WindowModel) Init() tea.Cmd {
	return nil
}

func (m WindowModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEscape {
			return m, tea.Quit
		}
	}

	switch m.window {
	case boardWindow:
		m.boardModel, cmd = m.boardModel.Update(msg)
		boardModel := m.boardModel.(*BoardModel)
		/**
		Here the part here is inferring into the child to find it's state and act on it, which makes too dependent on
		the child, maybe a better way to do it would be to have a simple active model that changed from the children
		*/
		if boardModel.phase == selected {
			m.playerModel = NewPlayerModel(boardModel.board.ID)
			cmd = m.playerModel.Init()
			m.window = playerWindow
		}
	case playerWindow:
		m.playerModel, cmd = m.playerModel.Update(msg)
		playerModel := m.playerModel.(*PlayerModel)
		boardModel := m.boardModel.(*BoardModel)
		/**
		Here the part here is inferring into the child to find it's state and act on it, which makes too dependent on
		the child, maybe a better way to do it would be to have a simple active model that changed from the children
		*/
		if playerModel.playerPhase == roleSelected {
			m.gameModel = NewGameModel(boardModel.board.ID, playerModel.player.ID)
			cmd = m.gameModel.Init()
			m.window = gameWindow
		}
	case gameWindow:
		m.gameModel, cmd = m.gameModel.Update(msg)
	default:
		panic("unhandled default case")
	}

	return m, cmd
}

func (m WindowModel) View() string {
	var content = ""
	switch m.window {
	case boardWindow:
		content = m.boardModel.View()
	case playerWindow:
		content = m.playerModel.View()
	case gameWindow:
		content = m.gameModel.View()
	default:
		panic("unhandled default case")
	}
	return content
}
