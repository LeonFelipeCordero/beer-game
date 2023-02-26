package neo4j

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories"
)

type Repository struct{}

func NewRepository() IRepository {
	return &Repository{}
}

func (r Repository) Save(ctx context.Context, value interface{}) error {
	session := repositories.GlobalSession(ctx)
	err := session.Save(ctx, value)

	if err != nil {
		return fmt.Errorf(
			fmt.Sprintf("Something went wrong creating new node"),
			err,
		)
	}

	err = session.Commit(ctx)

	if err != nil {
		return fmt.Errorf(
			fmt.Sprintf("Something executin transaction to create new node"),
			err,
		)
	}

	return nil
}

func (r Repository) Query(ctx context.Context, query string, values map[string]interface{}, target interface{}) error {
	session := repositories.GlobalSession(ctx)
	err := session.Query(context.Background(), query, values, target)

	if err != nil {
		return fmt.Errorf(
			fmt.Sprintf("Something went wrong getting data with query %s", query),
			err,
		)
	}

	return nil
}

func (r Repository) QueryRaw(ctx context.Context, query string, values map[string]interface{}) ([][]interface{}, error) {
	session := repositories.GlobalSession(ctx)

	result, _, err := session.QueryRaw(context.Background(), query, values)

	if err != nil {
		return result, fmt.Errorf(
			fmt.Sprintf("Something went wrong with raw query"),
			err,
		)
	}

	return result, nil
}
