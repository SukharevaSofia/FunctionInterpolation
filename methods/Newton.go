package methods

import "FunctionInterpolation/utils"

func NewtonPolynomial(data utils.XY, arg float64) (
	yValue float64,
	function func(x float64) float64,
	description string,
	estimationError float64) {

	n := data.GetLength()
	f := func(x float64) float64 {
		var sum float64 = 0.0
		for i := 1; i < n; i++ {
			var F float64 = 0.0
			for j := 0; j < i; j++ {
				var density float64 = 1.0
				for k := 0; k <= i; k++ {
					density *= data.X[j] - data.X[k]
				}
				F += data.Y[j] / density
			}
			for j := 0; j < i; j++ {
				F *= x - data.X[j]
			}
			sum += F
		}
		return sum
	}
	value := f(arg)
	errorEstimation := errorEstimation(data.X, f, arg)
	return value, f, "Полином Ньютона", errorEstimation
}
