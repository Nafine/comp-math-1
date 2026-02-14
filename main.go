package main

import (
	"fmt"

	"github.com/Nafine/comp-math-1/internal/matrix"
)

func main() {
	//p := tea.NewProgram(model.InitModel())
	//if _, err := p.Run(); err != nil {
	//	log.Fatal(err)
	//}

	if _, err := matrix.Solve(matrix.EquationSystem{
		Matrix: [][]float64{
			{1, 2, 3, 4, 11},
			{1, 2, 3, 11, 5},
			{1, 2, 12, 4, 5},
			{1, 13, 3, 4, 5},
			{14, 2, 3, 4, 5},
		},
		FreeTerms: []float64{
			1, 2, 3, 4, 5,
		},
	}); err != nil {
		fmt.Println(err)
	}
}
