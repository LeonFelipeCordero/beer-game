package repositories

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j/entities"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNeo4j(t *testing.T) {
	session := ConfigureDatabase()
	t.Run("Should get things from db", func(t *testing.T) {
		params := &entities.BoardNode{
			Name:      "test2",
			State:     "CREATED",
			Full:      false,
			Finished:  false,
			CreatedAt: time.Now().UTC(),
		}
		session.Begin(context.Background())

		var err = session.Save(context.Background(), params)

		session.Commit(context.Background())

		if err != nil {
			fmt.Printf("%e", err)
			panic("Should not fail")
		}

		boardNode := &entities.BoardNode{}
		query := "MATCH (b:BoardNode{name: $name}) RETURN count(b) as count"
		//err = session.Query(context.Background(), query, map[string]interface{}{
		//	"name": "test",
		//}, boardNode)
		result, _, err := session.QueryRaw(context.Background(), query, map[string]interface{}{
			"name": "test",
		})
		if err != nil {
			fmt.Printf("%e", err)
			fmt.Printf("%s", result)
			panic("Should not fail")
		}

		assert.Equal(t, boardNode.Name, params.Name)
		assert.Equal(t, boardNode.Full, params.Full)
		assert.Equal(t, boardNode.Finished, params.Finished)
		assert.Equal(t, boardNode.State, params.State)
		assert.Equal(t, boardNode.CreatedAt, params.CreatedAt)
	})
}
