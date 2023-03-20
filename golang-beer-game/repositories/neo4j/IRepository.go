package neo4j

import (
	"context"
)

type IRepository interface {
	Save(ctx context.Context, value interface{}) error
	SaveDepth(ctx context.Context, value interface{}) error
	Query(ctx context.Context, query string, values map[string]interface{}, target interface{}) error
	QueryRaw(ctx context.Context, query string, values map[string]interface{}) ([][]interface{}, error)
	LoadDepth(ctx context.Context, id int64, target interface{}) error
}
