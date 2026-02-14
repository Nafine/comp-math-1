package model

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type phase int

const (
	phaseSettings phase = iota
	phaseMatrix
	phaseSolution
)

const (
	n = iota
	eps
)

type Model struct {
	currentPhase phase

	inputs  []textinput.Model
	focused int

	matrix     [][]textinput.Model
	rows, cols int
	fRow, fCol int

	solution   []float64
	iterations int
	err        error
}

func InitModel() Model {
	inputs := make([]textinput.Model, 2)

	inputs[n] = textinput.New()
	inputs[n].Focus()

	inputs[n].Placeholder = "3"
	inputs[n].CharLimit = 2
	inputs[n].Width = 2
	inputs[n].Prompt = ""
	inputs[n].Validate = nValidator

	inputs[eps] = textinput.New()

	inputs[eps].Placeholder = "0.001"
	inputs[eps].CharLimit = 15
	inputs[eps].Width = 15
	inputs[eps].Prompt = ""
	inputs[eps].Validate = epsValidator

	return Model{
		currentPhase: phaseSettings,
		inputs:       inputs,
		focused:      0,
		err:          nil,
	}
}

func (m Model) Init() tea.Cmd {
	if len(os.Args) > 1 {
		return readFromFile(os.Args[1])
	}
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	if m.currentPhase == phaseSettings {
		return m.updateSettings(msg)
	}

	return m.updateMatrix(msg)
}

func (m Model) View() string {
	var s string

	if m.currentPhase == phaseSettings {
		s = m.renderSettings()
	} else {
		s = m.renderMatrix()
	}

	if m.err != nil {
		s += errorStyle.Render(fmt.Sprintf("\n\nError: %s\n", m.err.Error()))
	}
	return s
}
