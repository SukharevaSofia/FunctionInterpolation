package main

import (
	"FunctionInterpolation/IO"
	"FunctionInterpolation/utils"
	"fmt"
	"math"
)

// TODO: Многочлен Лагранжа,
// TODO: Многочлен Ньютона с конечными разностями
func main() {
	var eqn1 = utils.Equation{
		F:              func(x float64) float64 { return math.Pow(x, 2) },
		NameOfFunction: "y = x^2",
	}
	var eqn2 = utils.Equation{
		F:              func(x float64) float64 { return math.Pow(x, 3) },
		NameOfFunction: "y = x^3",
	}

	data := IO.UserInput(eqn1, eqn2)

	fmt.Println(data)

}
