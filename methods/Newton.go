package methods

import "FunctionInterpolation/utils"

func NewtonPolynomial(data utils.XY arg float64) float64 {
	n := len(data.X)
	p := data.Y
	for i := 1; i < n; i++ {
		x := 1.
		for j := 0; j < i; j++ {
			x *= arg - data.X[j]
		}
		p += f(i) * x
	}
	return p
}

func NewtonHalf(arg float64) float64 {
	h := DOTS[1][0] - DOTS[0][0]
	result := 0.
	if arg <= DOTS[int(len(DOTS)/2)][0] {

		targetX := DOTS[0][0]
		targetI := 0
		for i := 0; i <= len(DOTS)/2; i++ {
			if DOTS[i][0] < arg {
				targetX = DOTS[i][0]
				targetI = i
			} else {
				break
			}
		}
		t := (arg - targetX) / h
		result += Delta[0][targetI]
		num := t
		fact := 1.
		for i := 0; i < len(Delta); i++ {
			if len(Delta[i]) < targetI+2 {
				break
			}
			fact *= float64(i) + 1
			result += num * Delta[i+1][targetI] / fact
			num *= t - float64(i) - 1
		}
	} else {
		t := (arg - DOTS[len(DOTS)-1][0]) / h
		result += Delta[0][POINTS_COUNT-1]
		num := t
		fact := 1.
		for i := 0; i < len(Delta)-1; i++ {
			fact *= float64(i) + 1
			result += num * Delta[i+1][len(Delta[i+1])-1] / fact
			num *= t + float64(i) + 1
		}
	}
	return result
}

func f(k int) (float64) {
	k += 1
	result := 0.
	for i := 0; i < k; i++ {
		x := 1.
		for j := 0; j < k; j++ {
			if j != i {
				x *= data.X[i] - data.X[j]
			}
		}
		result += data.X[i][1]/x
	}
	return result
}

