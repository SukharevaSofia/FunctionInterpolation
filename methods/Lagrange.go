package methods

import (
	"FunctionInterpolation/utils"
	"math"
)

func LagrangeInterpolation(data utils.XY, arg float64) (yValue float64, description string, estimationError float64) {
	n := data.GetLength()
	f := func(x float64) float64 {
		sum := 0.0
		for i := 0; i < n; i++ {
			pv, pn := 1., 1.
			for j := 0; j < n; j++ {
				if i != j {
					pv *= x - data.X[j]
					pn *= data.X[i] - data.X[j]
				}
			}
			if pn >= 0.0 && pn <= 0e-9 {
				pn = 0.00001
			}
			sum += (pv * data.Y[i]) / pn
		}
		return sum
	}
	value := f(arg)
	errorEstim := errorEstimation(data.X, f, arg)
	return value, "Многочлен Лагранжа", errorEstim
}

func errorEstimation(xData []float64, f func(x float64) float64, arg float64) float64 {
	n := len(xData) - 1
	fn1 := f
	for i := 0; i < n+1; i++ {
		fn1 = utils.Derive(fn1)
	}
	f2 := func(x float64) float64 {
		p := 1.0
		for i := 0; i <= n; i++ {
			p *= x - xData[i]
		}
		return p
	}
	max := 0.0
	for _, element := range xData {
		max = math.Max(max, fn1(element))
	}
	return math.Abs(max * f2(arg) / float64(utils.Factorial(n+1)))
}
