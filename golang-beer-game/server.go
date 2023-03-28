package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/application/schedulers"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/adapters"
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
	streamers, eventChan := createEventBus()
	boardApiAdapter, playerApiAdapter, orderApiAdapter := createAdapters(eventChan)

	graphResolver := graph.Config{
		Resolvers: &resolver.Resolver{
			BoardApiAdapter:  boardApiAdapter,
			PlayerApiAdapter: playerApiAdapter,
			OrderApiAdapter:  orderApiAdapter,
			Streamers:        streamers,
		},
	}

	go events.EventHandler(streamers, eventChan)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graphResolver))

	srv.AddTransport(&transport.Websocket{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createAdapters(eventChan chan events.Event) (ports.IBoardApi, ports.IPlayerApi, ports.IOrderApi) {
	neo4j.ConfigureDatabase()
	neo4jRepository := neo4j.NewRepository()

	boardRepository := adapters2.NewBoardRepository(neo4jRepository)
	playerRepository := adapters2.NewPlayerRepository(neo4jRepository, boardRepository)
	orderRepository := adapters2.NewOrderRepository(neo4jRepository, playerRepository)

	boardService := application.NewBoardService(boardRepository, eventChan)
	playerService := application.NewPlayerService(playerRepository, boardService, eventChan)
	orderService := application.NewOrderService(orderRepository, playerService, boardService, eventChan)

	boardApiAdapter := adapters.NewBoardApiAdapter(boardService)
	playerApiAdapter := adapters.NewPlayerApiAdapter(playerService, boardService)
	orderApiAdapter := adapters.NewOrderApiAdapter(orderService, boardService)

	orderScheduler := schedulers.NewOrderScheduler(orderService)
	orderScheduler.Start()

	return boardApiAdapter, playerApiAdapter, orderApiAdapter
}

func createEventBus() (*events.Streamers, chan events.Event) {
	streamers := &events.Streamers{
		Streamers: map[string]events.Streamer{},
	}
	eventChan := make(chan events.Event)
	return streamers, eventChan
}
