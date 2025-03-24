package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/application/schedulers"
	"github.com/LeonFelipeCordero/golang-beer-game/graph"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/adapters"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/resolver"
	adapters2 "github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	storage "github.com/LeonFelipeCordero/golang-beer-game/repositories/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	streamers, eventChan := events.CreateEventBus()
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

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}).Handler)

	srv := handler.New(graph.NewExecutableSchema(graphResolver))
	configureServer(srv)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/graphql", srv)
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func createAdapters(eventChan chan events.Event) (adapters.BoardApiAdapter, adapters.PlayerApiAdapter, adapters.OrderApiAdapter) {
	ctx := context.Background()
	url := "postgres://beer_game:beer_game@localhost:5432/beer_game"
	connectionPool, _ := pgxpool.New(ctx, url)
	queries := storage.New(connectionPool)

	boardRepository := adapters2.NewBoardRepository(queries)
	playerRepository := adapters2.NewPlayerRepository(queries)
	orderRepository := adapters2.NewOrderRepository(queries)

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

func configureServer(srv *handler.Server) {
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
}
