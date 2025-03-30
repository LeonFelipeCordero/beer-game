package main

import (
	"LeonFelipeCorder/beer-game/clitool/app"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix(uuid.New().String())

	m := app.NewMainModel()
	//p := tea.NewProgram(m)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
