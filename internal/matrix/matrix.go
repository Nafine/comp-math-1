package matrix

type EquationSystem struct {
	ErrorMargin float64     `yaml:"errorMargin" env-required:"true"`
	Matrix      [][]float64 `yaml:"matrix" env-required:"true"`
	FreeTerms   []float64   `yaml:"freeTerms" env-required:"true"`
}

type EquationSystemSolution struct {
	Solved         bool
	Iterations     int
	SolutionVector []float64
	FinalMargin    float64
	MatrixNorm     float64
}
