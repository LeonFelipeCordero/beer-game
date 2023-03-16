package domain

import (
	"time"
)

type Board struct {
	Id        string
	Name      string
	State     State
	Full      bool
	Finished  bool
	CreatedAt time.Time
	Players   []Player
}

type State string

const (
	StateCreated  State = "CREATED"
	StateRunning  State = "RUNNING"
	StateFinished State = "FINISHED"
)

func (b *Board) HasRoleAvailable(role Role) bool {
	availableRoles := b.AvailableRoles()
	for _, availableRole := range availableRoles {
		if availableRole == role {
			return true
		}
	}
	return false
}

func (b *Board) AvailableRoles() []Role {
	availableRoles := []Role{RoleRetailer, RoleWholesaler, RoleFactory}
	for _, player := range b.Players {
		for index, role := range availableRoles {
			if role == player.Role {
				availableRoles = append(availableRoles[:index], availableRoles[index+1:]...)
			}
		}
	}
	return availableRoles
}

func (b *Board) GetPlayerByRole(role Role) Player {
	for _, player := range b.Players {
		if role == player.Role {
			return player
		}
	}
	panic("role doesn't exit in the board")
}

func (b *Board) Start() {
	b.Full = true
	b.State = StateRunning
}
