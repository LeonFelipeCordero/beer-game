package graph

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"log"
)

func mutate(ctx context.Context, query any, variables map[string]interface{}) ([]byte, error) {
	client := graphql.NewClient("http://localhost:8080/graphql", nil)
	response, err := client.MutateRaw(ctx, query, variables)
	if err != nil {
		log.Printf("Error querying GraphQL: %v\n", err)
		return nil, err
	}
	return response, nil
}

func query(ctx context.Context, query any, variables map[string]interface{}) ([]byte, error) {
	client := graphql.NewClient("http://localhost:8080/graphql", nil)
	response, err := client.QueryRaw(ctx, query, variables)
	if err != nil {
		log.Printf("Error querying GraphQL: %v\n", err)
		return nil, err
	}
	return response, nil
}

func subscribe(ctx context.Context, query interface{}, variables map[string]interface{}) (chan *[]byte, error) {
	client := graphql.NewSubscriptionClient("ws://localhost:8080/graphql")
	//defer client.Close()
	updates := make(chan *[]byte)
	go func() {
		client.Subscribe(query, variables, func(data []byte, err error) error {
			if err != nil {
				//fmt.Printf("Error sending subscribtion: %v\n", err.Error())
				log.Printf("Error subscribing sending subscription: %v\n", err.Error())
				return nil
			}
			if data == nil {
				return nil
			}

			updates <- &data

			return nil
		})
	}()

	go func() {
		err := client.Run()
		if err != nil {
			close(updates)
			log.Printf("Error running client: %v\n", err.Error())
		}
	}()
	return updates, nil
}
