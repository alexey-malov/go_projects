package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// The expression T(v) converts the value v to the type T
	var f float64 = math.Sqrt(float64(x*x + y*y))

	// implicit type conversioni is not allowed
	var z uint = uint(f)

	fmt.Println(x, y, z)
}
