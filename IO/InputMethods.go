package IO

import (
	"FunctionInterpolation/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserInput(eqn1, eqn2 utils.Equation) utils.XY {
	var inputType int
	fmt.Println(utils.INFO)
	fmt.Println(utils.CHOOSE_INPUT)
	for {
		fmt.Scan(&inputType)
		if !(inputType == 1 || inputType == 2 || inputType == 3) {
			fmt.Println(utils.INPUT_ERR)
			continue
		}
		break
	}
	switch inputType {
	case 1:
		return fileInput()
	case 2:
		return terminalInput()
	default:
		return functionInput(eqn1, eqn2)
	}
}

func fileInput() utils.XY {
	var numberOfElements int
	f, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	numberOfElements, _ = strconv.Atoi(line)
	dataX := make([]float64, numberOfElements)
	dataY := make([]float64, numberOfElements)
	scanner.Scan()
	lineX := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	lineY := strings.Split(scanner.Text(), " ")
	if (len(lineX) != numberOfElements) || (len(lineY) != numberOfElements) {
		fmt.Println(utils.INPUT_ERR)
		os.Exit(0)
	}
	for i := 0; i < numberOfElements; i++ {
		dataX[i], _ = strconv.ParseFloat(lineX[i], 64)
		dataY[i], _ = strconv.ParseFloat(lineY[i], 64)
	}

	return utils.XY{
		X: dataX,
		Y: dataY,
	}
}

func terminalInput() utils.XY {
	var numberOfElements int
	var lineX, lineY []string
	fmt.Print(utils.NUMBER_OF_DOTS)
	for {
		fmt.Scan(&numberOfElements)
		if numberOfElements >= 3 {
			break
		}
		fmt.Println(utils.NUMBER_OF_DOTS)
	}
	dataX := make([]float64, numberOfElements)
	dataY := make([]float64, numberOfElements)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(utils.INPUT_X)
		scanner.Scan()
		lineX = strings.Split(scanner.Text(), " ")
		fmt.Println(utils.INPUT_Y)
		scanner.Scan()
		lineY = strings.Split(scanner.Text(), " ")
		if (len(lineY) == numberOfElements) || (len(lineX) == numberOfElements) {
			break
		}
		fmt.Println(utils.INPUT_ERR)
	}
	for i := 0; i < numberOfElements; i++ {
		dataX[i], _ = strconv.ParseFloat(lineX[i], 64)
		dataY[i], _ = strconv.ParseFloat(lineY[i], 64)
	}

	return utils.XY{
		X: dataX,
		Y: dataY,
	}
}

func functionInput(eqn1, eqn2 utils.Equation) utils.XY {
	var inputValue, numberOfElements int
	var eqn utils.Equation
	fmt.Println(utils.CHOOSE_FUNC)
	fmt.Println("1: ", eqn1.NameOfFunction)
	fmt.Println("2: ", eqn2.NameOfFunction)

	for {
		fmt.Scan(&inputValue)
		if !(inputValue == 1 || inputValue == 2) {
			fmt.Println(utils.INPUT_ERR)
			continue
		}
		break
	}
	switch inputValue {
	case 1:
		eqn = eqn1
	default:
		eqn = eqn2
	}
	fmt.Println(utils.INPUT_INTERVAL)
	var x0, xn float64
	for {
		fmt.Scan(&x0, &xn)
		if x0 >= xn {
			fmt.Println(utils.INPUT_ERR)
			continue
		}
		break
	}

	fmt.Println(utils.NUMBER_OF_DOTS)
	for {
		fmt.Scan(&numberOfElements)
		if numberOfElements >= 3 {
			break
		}
		fmt.Println(utils.NUMBER_OF_DOTS)
	}
	dataX := []float64{}
	dataY := []float64{}
	h := (xn - x0) / float64(numberOfElements)

	for i := x0; i < xn; i += h {
		dataX = append(dataX, i)
		dataY = append(dataY, eqn.F(i))
	}

	return utils.XY{
		X: dataX,
		Y: dataY,
	}

}
