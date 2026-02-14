package model

import (
	"github.com/Nafine/comp-math-1/internal/matrix"

	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct{ err error }
type solutionMsg struct {
	solution   []float64
	iterations int
}

func readFromFile(filepath string) tea.Cmd {
	return func() tea.Msg {
		_, err := matrix.ReadCSV(filepath)

		if err != nil {
			return errMsg{err}
		}

		//solution, iterations := matrix.Solve(matrix.EquationSystem{})

		return solutionMsg{}
	}
}
