package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	step := 1.0
	for math.Abs(step) > 1e-10 {
		step = (z*z - x) / (2 * z)
		z -= step
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1000000), math.Sqrt(1000000))
}
