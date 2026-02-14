package model

import (
	"fmt"
	"strings"
)

func (m Model) renderSettings() string {
	s := fmt.Sprintf(
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

	return s
}

func (m Model) renderMatrix() string {
	var b strings.Builder
	b.WriteString("Fill up the matrix:\n\n")
	for r := 0; r < m.n; r++ {
		for c := 0; c < m.n; c++ {
			cellContent := m.matrix[r][c].View()
			b.WriteString(cellStyle.Render(cellContent))
		}
		b.WriteString("\n")
	}
	return b.String()
}
