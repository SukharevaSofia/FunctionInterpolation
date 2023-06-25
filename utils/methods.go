package utils

import "math/big"

const dx = 0.000001

func Derive(f func(x float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return (f(x+dx) - f(x)) / dx
	}
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

func XYtoMap(xy XY, precision float64) map[float64]float64 {
	xymap := map[float64]float64{}
	for i := 0; i < len(xy.X); i++ {
		xymap[Truncate(xy.X[i], precision)] = xy.Y[i]
	}
	return xymap
}

func Truncate(f float64, unit float64) float64 {
	bf := big.NewFloat(0).SetPrec(1000).SetFloat64(f)
	bu := big.NewFloat(0).SetPrec(1000).SetFloat64(unit)

	bf.Quo(bf, bu)

	// Truncate:
	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	f, _ = bf.Mul(bf, bu).Float64()
	return f
}
