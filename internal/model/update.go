package model

import (
	"fmt"
	"strconv"

	"github.com/Nafine/comp-math-1/internal/matrix"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateSettings(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				err := m.syncSettings()
				if err != nil {
					m.err = err
				} else {
					m.currentPhase = phaseMatrix
					m.err = nil
				}
			} else {
				m.nextInput()
			}
		case tea.KeyUp:
			m.prevInput()
		case tea.KeyDown:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
	case errMsg:
		m.err = msg.err
		return m, nil
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) updateMatrix(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp, tea.KeyDown, tea.KeyLeft, tea.KeyRight:
			m.extMatrixInputs[m.fRow][m.fCol].Blur()

			cur := m.extMatrixInputs[m.fRow][m.fCol]
			width := len(cur.Value())

			oldR, oldC := m.fRow, m.fCol

			if msg.Type == tea.KeyRight && cur.Position() == width {
				m.fCol++
				if m.fCol >= m.cols {
					m.fCol = 0
					m.fRow++
					if m.fRow >= m.rows {
						m.fRow = 0
					}
				}
			}
			if msg.Type == tea.KeyLeft && cur.Position() == 0 {
				m.fCol--
				if m.fCol < 0 {
					m.fCol = m.cols - 1
					m.fRow--
					if m.fRow < 0 {
						m.fRow = m.rows - 1
					}
				}
			}
			if msg.Type == tea.KeyDown {
				m.fRow++
				if m.fRow >= m.rows {
					m.fRow = 0
				}
			}
			if msg.Type == tea.KeyUp {
				m.fRow--
				if m.fRow < 0 {
					m.fRow = m.rows - 1
				}
			}

			m.extMatrixInputs[m.fRow][m.fCol].Focus()

			if oldR != m.fRow || oldC != m.fCol {
				if msg.Type == tea.KeyRight {
					m.extMatrixInputs[m.fRow][m.fCol].SetCursor(0)
				}
				if msg.Type == tea.KeyLeft {
					m.extMatrixInputs[m.fRow][m.fCol].SetCursor(len(m.extMatrixInputs[m.fRow][m.fCol].Value()))
				}

				return m, nil
			}
		case tea.KeyEnter:
			err := m.syncMatrix()
			if err != nil {
				m.err = err
			} else {
				m.currentPhase = phaseSolution
				m.err = nil
			}
		}
	}

	m.extMatrixInputs[m.fRow][m.fCol], cmd = m.extMatrixInputs[m.fRow][m.fCol].Update(msg)
	return m, cmd
}

func (m Model) updateSolution(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) updateError() (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m *Model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *Model) prevInput() {
	m.focused--

	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func (m *Model) syncSettings() error {
	if m.inputs[n].Err != nil {
		return m.inputs[n].Err
	}
	if m.inputs[eps].Err != nil {
		return m.inputs[eps].Err
	}
	valN, _ := strconv.Atoi(m.inputs[n].Value())
	m.rows = valN
	m.cols = valN + 1

	m.extMatrixInputs = make([][]textinput.Model, m.rows)
	for r := 0; r < m.rows; r++ {
		m.extMatrixInputs[r] = make([]textinput.Model, m.cols)
		for c := 0; c < m.cols; c++ {
			ti := textinput.New()
			ti.Prompt = ""
			ti.Width = 8
			ti.Placeholder = "0"
			ti.Validate = matrixCellValidator
			m.extMatrixInputs[r][c] = ti
		}
	}
	m.extMatrixInputs[0][0].Focus()
	m.fRow, m.fCol = 0, 0
	return nil
}

func (m *Model) syncMatrix() error {
	coeffMatrix := make([][]float64, m.rows)
	freeTerms := make([]float64, m.rows)
	errorMargin, err := strconv.ParseFloat(m.inputs[eps].Value(), 64)

	if err != nil {
		return err
	}

	for i, row := range m.extMatrixInputs {
		coeffMatrix[i] = make([]float64, m.rows)
		for j := 0; j < m.cols-1; j++ {
			val, err := strconv.ParseFloat(row[j].Value(), 64)

			if err != nil {
				return fmt.Errorf("invalid cell value: %s", err)
			}

			coeffMatrix[i][j] = val
		}

		freeTerm, err := strconv.ParseFloat(row[m.cols-1].Value(), 64)

		if err != nil {
			return fmt.Errorf("invalid free term value: %s", err)
		}

		freeTerms[i] = freeTerm
	}

	eqSystem := matrix.EquationSystem{
		ErrorMargin: errorMargin,
		Matrix:      coeffMatrix,
		FreeTerms:   freeTerms,
	}

	solutionSystem, err := matrix.Solve(&eqSystem)
	if err != nil {
		return err
	}

	m.solution = solutionSystem

	return nil
}
