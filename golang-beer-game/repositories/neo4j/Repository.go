package neo4j

import (
	"context"
)

type Repository struct{}

func NewRepository() IRepository {
	return Repository{}
}

func (r Repository) Save(ctx context.Context, value interface{}) error {
	session := GlobalSession()
	err := session.Save(ctx, value)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) SaveDepth(ctx context.Context, value interface{}) error {
	session := GlobalSession()
	err := session.SaveDepth(ctx, value, 2)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Query(ctx context.Context, query string, values map[string]interface{}, target interface{}) error {
	session := GlobalSession()
	err := session.Query(ctx, query, values, target)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) LoadDepth(ctx context.Context, id int64, target interface{}) error {
	session := GlobalSession()
	err := session.LoadDepth(ctx, target, id, 2)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) QueryRaw(ctx context.Context, query string, values map[string]interface{}) ([][]interface{}, error) {
	session := GlobalSession()

	result, _, err := session.QueryRaw(ctx, query, values)

	if err != nil {
		return nil, err
	}

	return result, nil
}
