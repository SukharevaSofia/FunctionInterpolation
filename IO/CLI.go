package IO

//
//import (
//	"FunctionInterpolation/utils"
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//	"unicode"
//)
//
//func ChooseInput() {
//	fmt.Println(utils.CHOOSE_INPUT)
//	var inputOption rune
//	for {
//		var inputStr string
//		_, _ = fmt.Scanln(&inputStr)
//		inputStr = inputStr + "\n"
//		inputOption = unicode.ToLower(rune(inputStr[0]))
//		if !('a' <= inputOption && inputOption <= 'c') {
//			fmt.Println(utils.INPUT_ERR)
//			continue
//		}
//		break
//	}
//
//	switch inputOption - 'a' {
//	case 0:
//		keyboardInput()
//	case 1:
//		а
//	default:
//		ф
//	}
//
//}
//
//func keyboardInput() {
//	fmt.Print(utils.NUMBER_OF_DOTS)
//	for {
//		fmt.Scan(&numberOfElements)
//		if numberOfElements <= 160 && numberOfElements >= 7 {
//			break
//		}
//		fmt.Println(NUMBER_OF_DOTS)
//	}
//	dataX = make([]float64, numberOfElements)
//	dataY = make([]float64, numberOfElements)
//	for {
//		scanner := bufio.NewScanner(os.Stdin)
//		fmt.Println(INPUT_X)
//		scanner.Scan()
//		lineX = strings.Split(scanner.Text(), " ")
//		fmt.Println(INPUT_Y)
//		scanner.Scan()
//		lineY = strings.Split(scanner.Text(), " ")
//		if (len(lineY) == numberOfElements) || (len(lineX) == numberOfElements) {
//			break
//		}
//		fmt.Println(INPUT_ERR)
//	}
//	for i := 0; i < numberOfElements; i++ {
//		dataX[i], _ = strconv.ParseFloat(lineX[i], 64)
//		dataY[i], _ = strconv.ParseFloat(lineY[i], 64)
//	}
//
//}
