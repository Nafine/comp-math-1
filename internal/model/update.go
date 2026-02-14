package model

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateSettings(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case solutionMsg:
		m.solution = msg.solution
		return m, tea.Quit
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
			m.matrix[m.fRow][m.fCol].Blur()

			cur := m.matrix[m.fRow][m.fCol]
			width := len(cur.Value())

			oldR, oldC := m.fRow, m.fCol

			if msg.Type == tea.KeyRight && cur.Position() == width {
				m.fCol++
				if m.fCol >= m.n {
					m.fCol = 0
					m.fRow++
					if m.fRow >= m.n {
						m.fRow = 0
					}
				}
			}
			if msg.Type == tea.KeyLeft && cur.Position() == 0 {
				m.fCol--
				if m.fCol < 0 {
					m.fCol = m.n - 1
					m.fRow--
					if m.fRow < 0 {
						m.fRow = m.n - 1
					}
				}
			}
			if msg.Type == tea.KeyDown {
				m.fRow++
				if m.fRow >= m.n {
					m.fRow = 0
				}
			}
			if msg.Type == tea.KeyUp {
				m.fRow--
				if m.fRow < 0 {
					m.fRow = m.n - 1
				}
			}

			m.matrix[m.fRow][m.fCol].Focus()

			if oldR != m.fRow || oldC != m.fCol {
				if msg.Type == tea.KeyRight {
					m.matrix[m.fRow][m.fCol].SetCursor(0)
				}
				if msg.Type == tea.KeyLeft {
					m.matrix[m.fRow][m.fCol].SetCursor(len(m.matrix[m.fRow][m.fCol].Value()))
				}

				return m, nil
			}
		case tea.KeyEnter:
			m.currentPhase = phaseSolution
		}
	}

	m.matrix[m.fRow][m.fCol], cmd = m.matrix[m.fRow][m.fCol].Update(msg)
	return m, cmd
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
	m.n = valN

	m.matrix = make([][]textinput.Model, m.n)
	for r := 0; r < m.n; r++ {
		m.matrix[r] = make([]textinput.Model, m.n+1)
		for c := 0; c <= m.n; c++ {
			ti := textinput.New()
			ti.Prompt = ""
			ti.Width = 8
			ti.Placeholder = "0"
			ti.Validate = matrixCellValidator
			m.matrix[r][c] = ti
		}
	}
	m.matrix[0][0].Focus()
	m.fRow, m.fCol = 0, 0
	return nil
}

func (m *Model) syncMatrix() error {
	//TODO
	return nil
}
