package main

import (
	"comp-math-1/internal/model"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.InitModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
