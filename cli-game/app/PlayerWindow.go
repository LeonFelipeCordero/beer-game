package app

import (
	"LeonFelipeCorder/beer-game/clitool/graph"
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

type roleItem struct {
	name        string
	description string
}

func (i roleItem) Title() string       { return i.name }
func (i roleItem) Description() string { return i.description }
func (i roleItem) FilterValue() string { return i.name }

type PlayerModel struct {
	roles             list.Model
	playerPhase       playerPhase
	player            graph.Player
	board             graph.Board
	boardGraphClient  graph.BoardGraphClient
	playerGraphClient graph.PlayerGraphClient
	boardUpdates      chan *graph.Board
}

type playerPhase int

const (
	roleSelection playerPhase = iota
	waiting
	roleSelected
	onPlayerError
)

func NewPlayerModel(boardId string) *PlayerModel {
	ctx := context.Background()
	boardClient := graph.NewBoardGraphClient()

	board, err := boardClient.GetBoardByID(ctx, boardId)
	if err != nil {
		log.Println("fatal: ", err)
		os.Exit(1)
	}

	var items []list.Item
	for _, availableRole := range board.AvailableRoles {
		switch *availableRole {
		case graph.RoleRetailer:
			items = append(items, roleItem{name: availableRole.String(), description: "Small local store RETAILER"})
		case graph.RoleWholesaler:
			items = append(items, roleItem{name: availableRole.String(), description: "Wholesaler distributor for a zone"})
		case graph.RoleFactory:
			items = append(items, roleItem{name: availableRole.String(), description: "Small brewery with limited production"})
		}
	}

	updatesChannel, err := boardClient.SubscribeToBoard(ctx, boardId)
	if err != nil {
		log.Fatalf("fatal, can not get board updates: %s\n", err.Error())
	}

	m := PlayerModel{
		roles:             list.New(items, list.NewDefaultDelegate(), 150, 50),
		playerPhase:       roleSelection,
		board:             *board,
		boardGraphClient:  graph.NewBoardGraphClient(),
		playerGraphClient: graph.NewPlayerGraphClient(),
		boardUpdates:      updatesChannel,
	}
	m.roles.Title = "Choose a Role"
	m.roles.SetShowFilter(false)
	m.roles.SetShowHelp(false)
	m.roles.SetShowPagination(false)
	m.roles.SetShowStatusBar(false)

	return &m
}

func (m *PlayerModel) Init() tea.Cmd {
	return tea.Batch(
		waitForBoardUpdate(m.boardUpdates),
	)
}

func (m *PlayerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case tea.KeyMsg:
		cmd = m.phaseHandler(msg.(tea.KeyMsg))
	case *graph.Board:
		m.board = *msg.(*graph.Board)
		if m.playerPhase == roleSelection {
			displayedRoles := m.roles.Items()
			var indexesToDelete []int
			for index, displayedRole := range displayedRoles {
				isContained := false
				for _, availableRole := range m.board.AvailableRoles {
					if availableRole.String() == displayedRole.(roleItem).name {
						isContained = true
						break
					}
				}
				if !isContained {
					indexesToDelete = append(indexesToDelete, index)
				}
			}
			for _, index := range indexesToDelete {
				m.roles.RemoveItem(index)
			}

		} else if m.playerPhase == waiting {
			if m.board.Full && m.board.State == graph.BoardStateRunning {
				m.playerPhase = roleSelected
			}
		}

		cmd = waitForBoardUpdate(m.boardUpdates)
	}

	return m, cmd
}

func (m *PlayerModel) phaseHandler(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	switch {
	case m.playerPhase == roleSelection:
		if msg.String() == "enter" {
			ctx := context.Background()
			selectedItem := m.roles.SelectedItem()
			player, err := m.playerGraphClient.AddPlayer(ctx, m.board.ID, selectedItem.(roleItem).name)
			if err != nil {
				m.playerPhase = onPlayerError
				log.Printf("Something went wrong %s\n", err.Error())
				return nil
			}
			m.player = *player

			m.playerPhase = waiting
		}
		m.roles, cmd = m.roles.Update(msg)
	case m.playerPhase == waiting:
		if msg.String() == "enter" {
			m.playerPhase = roleSelection
		}
	case m.playerPhase == onPlayerError:
		if msg.String() == "enter" {
			m.playerPhase = roleSelection
		}
	}

	return cmd
}

func (m *PlayerModel) View() string {
	var content string
	switch m.playerPhase {
	case roleSelection:
		content = m.roles.View()
	case waiting:
		content = fmt.Sprintf("Waiting for gmae to start: %s - %s", m.player.ID, m.player.Role)
	case roleSelected:
		content = fmt.Sprintf("All set...")
	case onPlayerError:
		content = "something went wrong, press 'esc' to quit or enter to go back"
	}

	return docStyle.Render(content)
}
