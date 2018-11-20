package main

import "fmt"

func main() {
	fmt.Print("Enter temperature in Fahrengeit degrees: ")

	var fahr float64
	if _, err := fmt.Scanf("%f", &fahr); err != nil {
		panic(err)
	}

	var celsius = (fahr - 32) * 5.0 / 9.0

	fmt.Printf("%.1f in Celsius is: %.1f\n", fahr, celsius)
}
