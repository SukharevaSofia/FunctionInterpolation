package utils

type Equation struct {
	F              func(x float64) float64
	NameOfFunction string
}

type XY struct {
	X []float64
	Y []float64
}
