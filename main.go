package main

import (
	"FunctionInterpolation/IO"
	"FunctionInterpolation/methods"
	"FunctionInterpolation/utils"
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"math"
	"os"
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
	fmt.Println("XXX: ", data.X, "YYY: ", data.Y)
	fmt.Println(utils.INPUT_X0)
	var argument float64
	fmt.Scan(&argument)
	newtonResult, newtonName, newtonErr := methods.NewtonPolynomial(data, argument)
	lagrResult, lagrName, largErr := methods.LagrangeInterpolation(data, argument)
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 2,
				},
				Name:    "Узлы интерполяции",
				XValues: data.X,
				YValues: data.Y,
			}, chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 3,
				},
				Name:    newtonName,
				XValues: []float64{argument},
				YValues: []float64{newtonResult},
			}, chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 3,
				},
				Name:    lagrName,
				XValues: []float64{argument},
				YValues: []float64{lagrResult},
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	picture, _ := os.Create("graph.png")
	graph.Render(chart.PNG, picture)
	picture.Close()
	fmt.Println("Найденые значения: Ньютон - ", newtonResult, "; Лагранж - ", lagrResult)
	fmt.Println("Оценки погрешностей: Ньютон - ", newtonErr, "; Лагранж - ", largErr)
}
