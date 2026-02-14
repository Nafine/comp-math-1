package matrix

import (
	"math"
)

const maxIterations = 1000

func Solve(s *EquationSystem) (EquationSystemSolution, error) {
	if err := s.toDiagonalDominant(); err != nil {
		return EquationSystemSolution{}, err
	}

	s.transform()

	return s.calc(), nil
}

func (s *EquationSystem) calc() EquationSystemSolution {
	prev := make([]float64, len(s.Matrix))
	cur := make([]float64, len(s.Matrix))

	copy(prev, s.FreeTerms)

	iterations := 0

	for iterations < maxIterations {
		iterations++
		for i := 0; i < len(s.Matrix); i++ {
			cur[i] = s.FreeTerms[i] + s.sum(prev, cur, i)
		}

		if absDelta(prev, cur) < s.ErrorMargin {
			break
		}

		copy(prev, cur)
	}

	return EquationSystemSolution{
		Solved:         true,
		Iterations:     iterations,
		SolutionVector: cur,
	}
}

func (s *EquationSystem) sum(prev, cur []float64, i int) float64 {
	sum := 0.0

	for j := 0; j < i; j++ {
		sum += s.Matrix[i][j] * cur[j]
	}

	for j := i + 1; j < len(prev); j++ {
		sum += s.Matrix[i][j] * prev[j]
	}

	return sum
}

func absDelta(prev, cur []float64) float64 {
	maxDelta := math.Abs(cur[0] - prev[0])

	for i := 1; i < len(cur); i++ {
		delta := math.Abs(cur[i] - prev[i])
		if delta > maxDelta {
			maxDelta = delta
		}
	}

	return maxDelta
}
