package main

import "fmt"

func main() {
	fmt.Print("Enter length in feet: ")

	var feet float64
	if _, err := fmt.Scanf("%f", &feet); err != nil {
		panic(err)
	}

	var meters = feet * 0.3048
	fmt.Println(feet, "ft is", meters, "m")
}
