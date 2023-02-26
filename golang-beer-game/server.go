package main

import (
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/adapters"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories"
	adapters2 "github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LeonFelipeCordero/golang-beer-game/graph"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/resolver"
)

const defaultPort = "8080"

func main() {
	repositories.ConfigureDatabase()
	neo4jRepository := neo4j.NewRepository()
	boardRepository := adapters2.NewBoardRepository(neo4jRepository)
	boardService := application.NewBoardService(boardRepository)
	boardApiAdapter := adapters.NewBoardApiAdapter(boardService)

	gogmResolver := graph.Config{
		Resolvers: &resolver.Resolver{
			BoardApiAdapter: boardApiAdapter,
		},
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(gogmResolver))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
