package matrix

import (
	"fmt"
	"math"
)

func (s *EquationSystem) toDiagonalDominant() error {
	if s.isDiagonalDominant() {
		return nil
	}

	maxIds, err := s.findMaxIds()

	if err != nil {
		return err
	}

	//check if multiple rows dominant on one index
	for i := 0; i < len(s.Matrix); i++ {
		for j := i + 1; j < len(s.Matrix); j++ {
			if maxIds[i] == maxIds[j] {
				return fmt.Errorf("impossible to reach diagonal dominance, " +
					"multiple rows are dominant on same index")
			}
		}
	}

	s.swapRows(maxIds)

	if s.isDiagonalDominant() {
		return nil
	}

	return fmt.Errorf("impossible to reach diagonal dominance")
}

func (s *EquationSystem) isDiagonalDominant() bool {
	hasStrictlyGreater := false

	for i := 0; i < len(s.Matrix); i++ {
		var lineSum float64
		var diagonalFactor float64

		for j := 0; j < len(s.Matrix[i]); j++ {
			absVal := math.Abs(s.Matrix[i][j])
			if i == j {
				diagonalFactor = absVal
			} else {
				lineSum += absVal
			}
		}

		if diagonalFactor < lineSum {
			return false
		}

		if diagonalFactor > lineSum {
			hasStrictlyGreater = true
		}
	}

	return hasStrictlyGreater
}

func (s *EquationSystem) findMaxIds() ([]int, error) {
	maxIds := make([]int, len(s.Matrix))

	for i := range maxIds {
		maxIdx, err := maxRowValIdx(s.Matrix[i])

		if err != nil {
			return nil, err
		}

		maxIds[i] = maxIdx
	}

	return maxIds, nil
}

func maxRowValIdx(row []float64) (int, error) {
	maxVal := math.Abs(row[0])
	maxIdx := 0

	for i := 1; i < len(row); i++ {
		absVal := math.Abs(row[i])
		if absVal > maxVal {
			maxVal = absVal
			maxIdx = i
		}
	}

	if maxVal < 1e-15 {
		return 0, fmt.Errorf("zero row detected")
	}

	return maxIdx, nil
}

func (s *EquationSystem) swapRows(maxRowValIds []int) {
	n := len(s.Matrix)
	newMatrix := make([][]float64, n)
	newFreeTerms := make([]float64, n)

	for oldIdx, newIdx := range maxRowValIds {
		newMatrix[newIdx] = s.Matrix[oldIdx]
		newFreeTerms[newIdx] = s.FreeTerms[oldIdx]
	}

	s.Matrix = newMatrix
	s.FreeTerms = newFreeTerms
}
