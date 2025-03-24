package application

import (
	testingutil "github.com/LeonFelipeCordero/golang-beer-game/testing"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testingutil.Setup()

	code := m.Run()

	os.Exit(code)
}
