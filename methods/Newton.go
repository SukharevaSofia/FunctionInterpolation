package methods

import (
	"FunctionInterpolation/utils"
	"fmt"
)

func NewtonPolynomial(data utils.XY, arg float64) (yValue float64, description string, estimationError float64) {
	n := data.GetLength()
	f := func(x float64) (float64, [][]float64) {
		table := make([][]float64, n)
		for i := range table {
			table[i] = make([]float64, n)
		}

		for index, value := range data.Y {
			table[index][0] = value
		}
		for j := 1; j < n; j++ {
			for i := 0; i < n-j; i++ {
				table[i][j] = table[i+1][j-1] - table[i][j-1]
			}
		}

		//AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

		h := data.X[1] - data.X[0]
		xMid := (data.X[0] + data.X[n-1]) / 2
		if x > xMid {
			t := (x - data.X[n-1]) / h
			result := table[n-1][0]
			for i := 1; i < n; i++ {
				compT := 1.0
				currT := t
				for j := 0; j < i; j++ {
					compT *= currT
					currT += 1
				}
				result += compT * table[n-i-1][i] / float64(utils.Factorial(i))
			}
			return result, table
		} else {
			leftIndex := -1
			for i := 0; i < n-1; i++ {
				if data.X[i] < x && x < data.X[i+1] {
					leftIndex = i
					break
				} else if x == data.X[i] {
					return data.Y[i], table
				} else if x == data.X[i+1] {
					return data.Y[i+1], table
				}
			}
			if leftIndex == -1 {
				if x > data.X[n-1] {
					leftIndex = n - 1
				} else {
					leftIndex = 0
				}
			}
			t := (x - data.X[leftIndex]) / h
			result := table[leftIndex][0]
			for i := 1; i < n-leftIndex; i++ {
				compT := 1.0
				currT := t
				for j := 0; j < i; j++ {
					compT *= currT
					currT--
				}
				result += compT * table[leftIndex][i] / float64(utils.Factorial(i))
			}
			return result, table
		}

	}
	fWrapValue := func(x float64) float64 {
		val, _ := f(x)
		return val
	}
	value := fWrapValue(arg)
	errorEst := errorEstimation(data.X, fWrapValue, arg)
	_, table := f(arg)
	fmt.Println("Таблица конечных разностей")
	for i, list := range table {
		for j, element := range list {
			if n-i > j {
				fmt.Printf("%f ", element)
			} else {
				fmt.Printf("")
			}
		}
		fmt.Println("")
	}
	return value, "Полином Ньютона", errorEst
}
