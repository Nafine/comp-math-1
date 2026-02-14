package main

import (
	"log"

	"github.com/Nafine/comp-math-1/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.InitModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
