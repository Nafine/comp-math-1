package matrix

func (s *EquationSystem) transform() {
	for i := range s.Matrix {
		if s.Matrix[i][i] == 0 {
			continue
		}

		term := s.Matrix[i][i]

		for j := range s.Matrix[i] {
			if i != j {
				s.Matrix[i][j] *= -1
				s.Matrix[i][j] /= term
			} else {
				s.Matrix[i][j] = 0
			}
		}

		s.FreeTerms[i] /= term
	}
}
