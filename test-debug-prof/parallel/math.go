package math

import "time"

func Mul(x, y float64) float64 {
	time.Sleep(1 * time.Second)
	return x * y
}

func Div(x, y float64) float64 {
	time.Sleep(1 * time.Second)
	return x / y
}
