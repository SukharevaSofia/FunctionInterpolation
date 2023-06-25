package methods

import (
	"FunctionInterpolation/utils"
	"fmt"
)

func NewtonPolynomial(data utils.XY, arg float64) (
	yValue float64,
	function func(x float64) float64,
	description string,
	estimationError float64) {

	n := data.GetLength()
	f := func(x float64) float64 {
		table := make([][]float64, n)
		for i := range table {
			table[i] = make([]float64, n)
		}

		for index, value := range data.Y {
			table[index][0] = value
		}
		for j := 1; j < n; j++ {
			for i := 0; i < n-j; i++ {
				table[i][j] = (table[i+1][j-1] - table[i][j-1]) / (data.X[i+j] - data.X[i])
			}
		}

		// Table
		fmt.Println("Таблица конечных разностей")
		for _, list := range table {
			for _, element := range list {
				fmt.Printf("%f ", element)
			}
			fmt.Println("")
		}

		var sum = 0.0
		for i := 0; i < n; i++ {
			var diff = table[0][i]
			for j := 0; j < i; j++ {
				diff *= x - data.X[j]
			}
		}
		return sum
	}
	value := f(arg)
	errorEstimation := errorEstimation(data.X, f, arg)
	return value, f, "Полином Ньютона", errorEstimation
}
