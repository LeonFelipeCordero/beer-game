package app

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

var tickerDocStyle = lipgloss.NewStyle().Margin(1, 2)

type responseMsg struct {
	Msg string
}

type myListItem struct {
	name        string
	description string
}

func (i myListItem) Title() string       { return i.name }
func (i myListItem) Description() string { return i.description }
func (i myListItem) FilterValue() string { return i.name }

func listenForActivity(sub chan responseMsg) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)+100))
			sub <- responseMsg{
				Msg: uuid.New().String(),
			}
		}
	}
}

func waitForActivity(sub chan responseMsg) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

type TickerModel struct {
	sub         chan responseMsg
	responses   int
	lastMessage string
	spinner     spinner.Model
	options     list.Model
	quitting    bool
}

func CreateNewTicker() TickerModel {
	items := []list.Item{
		myListItem{name: "TEST 1", description: "This should be first removed"},
		myListItem{name: "TEST 2", description: "This should be first removed"},
		myListItem{name: "TEST 3", description: "This should be first removed"},
	}

	return TickerModel{
		sub:     make(chan responseMsg),
		spinner: spinner.New(),
		options: list.New(items, list.NewDefaultDelegate(), 150, 18),
	}
}

func (m TickerModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		listenForActivity(m.sub),
		waitForActivity(m.sub),
	)
}

func (m TickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case tea.KeyMsg:
		if msg.(tea.KeyMsg).String() == "q" {
			m.quitting = true
			return m, tea.Quit
		}
		m.options, cmd = m.options.Update(msg)
	case responseMsg:
		m.responses++
		m.lastMessage = msg.(responseMsg).Msg
		m.options.RemoveItem(0)
		cmd = waitForActivity(m.sub)
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
	}
	return m, cmd
}

func (m TickerModel) View() string {
	listView := m.options.View()
	events := fmt.Sprintf("\n %s Events received: %d\n\n", m.spinner.View(), m.responses)
	lastEvent := fmt.Sprintf("\n last event id %s\n", m.lastMessage)
	s := fmt.Sprintf("%s\n%s%s", listView, events, lastEvent)
	if m.quitting {
		s += "\n"
	}

	return tickerDocStyle.Render(s)
}
