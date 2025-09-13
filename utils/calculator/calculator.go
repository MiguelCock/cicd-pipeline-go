package calculator

import "errors"

func Sum(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mult(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("error: Division by zero")
	}
	return a / b, nil
}	
