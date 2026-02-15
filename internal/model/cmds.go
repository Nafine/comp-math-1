package model

import (
	"github.com/Nafine/comp-math-1/internal/matrix"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct{ err error }

func readFromFile(filepath string) tea.Cmd {
	return func() tea.Msg {
		var es matrix.EquationSystem

		err := matrix.ReadYaml(filepath, &es)

		if err != nil {
			return errMsg{err: err}
		}

		solution, err := matrix.Solve(&es)

		if err != nil {
			return errMsg{err: err}
		}

		return solution
	}
}
