package matrix

type EquationSystem struct {
	ErrorMargin float64
	Matrix      [][]float64
	FreeTerms   []float64
}

type EquationSystemSolution struct {
	Solved         bool
	Iterations     int
	SolutionVector []float64
	FinalMargin    float64
	MatrixNorm     float64
}
