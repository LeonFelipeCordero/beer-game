package testingutil

import (
	"context"
	"fmt"
	storage "github.com/LeonFelipeCordero/golang-beer-game/repositories/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"path/filepath"
)

var (
	postgresContainer testcontainers.Container
	connection        string
)

func Setup() {
	fmt.Println("Setting up testing environment...")

	startTestContainer()
	connectionUrl := GetPostgresConnection()

	root, err := findProjectRoot()
	if err != nil {
		panic(err)
	}
	migrationsPath := filepath.Join("file://", root, "db", "migrations")

	//"file://../../db/migrations",
	migrateInstance, err := migrate.New(
		migrationsPath,
		connectionUrl,
	)
	if err != nil {
		fmt.Printf("Impossible to stablish connection for migration during tetsting: %s", err.Error())
		os.Exit(0)
	}

	err = migrateInstance.Up()
	if err != nil && err.Error() != "no change" {
		fmt.Printf("Impossible to run migrations during testing: %s", err.Error())
		os.Exit(0)
	}
}

func Clean(ctx context.Context, queries *storage.Queries) {
	queries.DeleteAllOrders(ctx)
	queries.DeleteAllPlayers(ctx)
	queries.DeleteAllBoards(ctx)
}

func SetupDatabaseConnection(ctx context.Context) *storage.Queries {
	conn, err := pgx.Connect(ctx, GetPostgresConnection())

	if err != nil {
		panic(fmt.Sprintf("Impossible to establihs database connectioin: %s", err.Error()))
	}

	queries := storage.New(conn)

	return queries
}

func GetPostgresConnection() string {
	return fmt.Sprintf("postgres://%s", connection)
}

func startTestContainer() {
	user := "beer_game"
	password := "beer_game"
	database := "beer_game"
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     user,
			"POSTGRES_PASSWORD": password,
			"POSTGRES_DB":       database,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
		Name:       "postgres-testcontainers",
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            true,
	})

	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}

	postgresContainer = container

	host, _ := container.Host(ctx)
	port, _ := container.MappedPort(ctx, "5432")

	connection = fmt.Sprintf(
		"%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port.Port(), database,
	)
	fmt.Println(fmt.Sprintf("Setting up connection: %s", connection))
}

func findProjectRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, nil // Found root directory containing go.mod
		}

		parent := filepath.Dir(wd)
		if parent == wd { // Reached system root
			return "", os.ErrNotExist
		}

		wd = parent
	}
}
