package main

import "errors"

func sum(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mult(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("error: Division by zero")
	}
	return a / b, nil
}	
