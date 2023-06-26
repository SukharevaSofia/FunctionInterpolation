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
	if data.GetLength() > 18 {
		fmt.Println(utils.TOO_MUCH_DOTS)
		os.Exit(0)
	}
	if true {
		// Previous precision checking routine
		h := math.Round(data.X[1]*1000) - math.Round(data.X[0]*1000)
		for i := 1; i < len(data.X); i++ {
			if math.Round(data.X[i]*1000)-math.Round(data.X[i-1]*1000) != h {
				fmt.Println(math.Round(data.X[i]*1000)-math.Round(data.X[i-1]*1000), h)
				fmt.Println(utils.BAD_DOTS)
				os.Exit(0)
			}
		}
	} else {
		precision := 0.1
		for i := 1; i < data.GetLength(); i++ {
			if math.Abs(data.X[i]-data.X[i-1]) > precision {
				fmt.Println(utils.BAD_DOTS)
				os.Exit(0)
			}
		}
	}

	fmt.Println(utils.INPUT_X0)
	var argument float64
	fmt.Scan(&argument)
	if (argument > data.X[len(data.X)-1]) || (argument < data.X[0]) {
		fmt.Println(utils.BAD_DOT)
		os.Exit(0)
	}
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
					DotWidth: 5,
				},
				Name:    newtonName,
				XValues: []float64{argument},
				YValues: []float64{newtonResult},
			}, chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 2,
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
