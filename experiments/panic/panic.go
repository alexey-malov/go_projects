package main

import (
	"fmt"
	"math"
)

func MySqrt(v float64) float64 {
	if v < 0 {
		panic("value is less than 0")
	} else {
		return math.Sqrt(v)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Err:", err)
		}
	}()
	fmt.Println("Sqrt is", MySqrt(-3))
}
