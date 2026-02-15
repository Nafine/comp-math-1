package model

import (
	"fmt"
	"strings"
)

func (m Model) renderSettings() string {
	return fmt.Sprintf(
		`
%s
%s

%s
%s
`,
		inputStyle.Width(11).Render("Matrix Size"),
		m.inputs[n].View(),
		inputStyle.Width(7).Render("Epsilon"),
		m.inputs[eps].View(),
	)
}

func (m Model) renderMatrix() string {
	var b strings.Builder
	b.WriteString("Fill up the extMatrixInputs:\n\n")

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols-1; c++ {
			cellContent := m.extMatrixInputs[r][c].View()
			b.WriteString(cellStyle.Render(cellContent))
		}
		//Free terms of an equation
		b.WriteString(" | ")
		b.WriteString(cellStyle.Render(m.extMatrixInputs[r][m.cols-1].View()))
		b.WriteString("\n")
	}
	return b.String()
}

func (m Model) renderSolution() string {
	return fmt.Sprintf(
		`
%s
%f

%s
%f

%s
%d

%s
%f
`,
		inputStyle.Width(11).Render("Matrix norm"),
		m.solution.MatrixNorm,
		inputStyle.Width(15).Render("Solution vector"),
		m.solution.SolutionVector,
		inputStyle.Width(10).Render("Iterations"),
		m.solution.Iterations,
		inputStyle.Width(12).Render("Final margin"),
		m.solution.FinalMargin,
	)
}
