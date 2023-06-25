package utils

type Equation struct {
	F              func(x float64) float64
	NameOfFunction string
}

type XY struct {
	X []float64
	Y []float64
}

func (data XY) GetLength() int {
	if len(data.X) != len(data.Y) {
		panic("X length does not equal Y length")
	}

	return len(data.X) - 1
}
