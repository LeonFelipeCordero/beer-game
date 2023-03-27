//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BoardApiAdapter  ports.IBoardApi
	PlayerApiAdapter ports.IPlayerApi
	OrderApiAdapter  ports.IOrderApi
	Streamers        *events.Streamers
}
