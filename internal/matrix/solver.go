package matrix

func Solve(s EquationSystem) (EquationSystemSolution, error) {
	if err := s.toDiagonalDominant(); err != nil {
		return EquationSystemSolution{}, err
	}

	s.transform()

	return EquationSystemSolution{}, nil
}
