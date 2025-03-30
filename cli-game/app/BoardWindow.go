package app

import (
	"LeonFelipeCorder/beer-game/clitool/graph"
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type action int

const (
	create action = iota
	find
)

type item struct {
	title, description string
	action             action
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type BoardModel struct {
	options          list.Model
	phase            Phase
	board            graph.Board
	boardInput       BoardInput
	boardGraphClient graph.BoardGraphClient
}
type BoardInput struct {
	NameInput textinput.Model
	action    action
	error     string
}
type Phase int

const (
	selection Phase = iota
	typing
	selected
	onError
)

func NewBoardModel() *BoardModel {
	ti := textinput.New()
	ti.Focus()

	items := []list.Item{
		item{title: "Create a new board", description: "You can invite other players to join the board", action: create},
		item{title: "Choose an existing board", description: "Play in your friends boards", action: find},
	}

	m := BoardModel{
		options:          list.New(items, list.NewDefaultDelegate(), 0, 0),
		phase:            selection,
		boardInput:       BoardInput{NameInput: ti},
		boardGraphClient: graph.NewBoardGraphClient(),
	}
	m.options.Title = "Choose an option"
	m.options.SetShowFilter(false)
	m.options.SetShowHelp(false)
	m.options.SetShowPagination(false)
	m.options.SetShowStatusBar(false)

	return &m
}

func (m *BoardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *BoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var model tea.Model = m
	switch msg := msg.(type) {
	case tea.KeyMsg:
		model, cmd = m.phaseHandler(msg)
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.options.SetSize(msg.Width-h, msg.Height-v)
	}

	return model, cmd
}

func (m *BoardModel) phaseHandler(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var model tea.Model = m
	switch {
	case m.phase == selection:
		if msg.String() == "enter" {
			selectedItem := m.options.SelectedItem()
			if selectedItem.(item).action == create {
				m.boardInput.action = create
			} else {
				m.boardInput.action = find
			}
			m.phase = typing
		}
		m.options, cmd = m.options.Update(msg)
	case m.phase == typing:
		m.boardInput.NameInput, cmd = m.boardInput.NameInput.Update(msg)
		if msg.String() == "enter" {
			ctx := context.Background()
			m.board.Name = m.boardInput.NameInput.Value()
			if m.boardInput.action == create {
				board, err := m.boardGraphClient.CreateBoard(ctx, m.board.Name)
				if err != nil {
					m.phase = onError
					m.boardInput.error = err.Error()
					return model, cmd
				}
				m.board = *board
			} else {
				board, err := m.boardGraphClient.GetBoardByName(ctx, m.board.Name)
				if err != nil {
					m.phase = onError
					m.boardInput.error = err.Error()
					// how to not return here
					return model, cmd
				}
				m.board = *board
			}
			m.boardInput.NameInput.Reset()
			m.phase = selected

			// in case we want how main model control the flow we can potentially change to this approach where the children
			// return a different view
			//model = NewPlayerModel(m.board.Id)
			//cmd = model.Init()
		}
	case m.phase == selected:
	case m.phase == onError:
		if msg.String() == "enter" {
			m.phase = typing
			m.board = graph.Board{}
		}
	}

	return model, cmd
}

func (m *BoardModel) View() string {
	var content string
	switch m.phase {
	case selection:
		content = m.options.View()
	case typing:
		boardNameInput := m.boardInput.NameInput.Value()
		if m.boardInput.action == create {
			content = fmt.Sprintf("Enter new board name:\n %s", boardNameInput)
		} else {
			content = fmt.Sprintf("Enter board name to search:\n %s", boardNameInput)
		}
	case selected:
		content = fmt.Sprintf("Your board: %s - %s\nPress 'esc' to quit or confirm with enter", m.board.Name, m.board.ID)
	case onError:
		content = m.boardInput.error + "\n press 'esc' to quit or enter to go back"
	}

	return docStyle.Render(content)
}
