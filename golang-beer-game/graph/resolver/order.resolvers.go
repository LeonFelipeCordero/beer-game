package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"

	"github.com/LeonFelipeCordero/golang-beer-game/graph"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, receiverID *string) (*model.Order, error) {
	return r.OrderApiAdapter.CreateOrder(ctx, *receiverID)
}

// DeliverOrder is the resolver for the deliverOrder field.
func (r *mutationResolver) DeliverOrder(ctx context.Context, orderID *string, amount *int64) (*model.Response, error) {
	return r.OrderApiAdapter.DeliverOrder(ctx, *orderID, *amount)
}

// Sender is the resolver for the sender field.
func (r *orderResolver) Sender(ctx context.Context, obj *model.Order) (*model.Player, error) {
	return r.PlayerApiAdapter.Get(ctx, obj.SenderId)
}

// Receiver is the resolver for the receiver field.
func (r *orderResolver) Receiver(ctx context.Context, obj *model.Order) (*model.Player, error) {
	return r.PlayerApiAdapter.Get(ctx, obj.ReceiverId)
}

// Board is the resolver for the board field.
func (r *orderResolver) Board(ctx context.Context, obj *model.Order) (*model.Board, error) {
	return r.BoardApiAdapter.GetByPlayer(ctx, obj.SenderId)
}

// NewOrder is the resolver for the newOrder field.
func (r *subscriptionResolver) NewOrder(ctx context.Context, playerID *string) (<-chan *model.Order, error) {
	return r.OrderApiAdapter.NewOrderSubscription(ctx, *playerID, r.Streamers)
}

// OrderDelivery is the resolver for the orderDelivery field.
func (r *subscriptionResolver) OrderDelivery(ctx context.Context, playerID *string) (<-chan *model.Order, error) {
	return r.OrderApiAdapter.OrderDeliveredSubscription(ctx, *playerID, r.Streamers)
}

// Order returns graph.OrderResolver implementation.
func (r *Resolver) Order() graph.OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }
