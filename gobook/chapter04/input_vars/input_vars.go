package main

import "fmt"

func main() {
	fmt.Print("Enter a number ")

	var number float64
	_, err := fmt.Scanf("%f", &number)
	if err != nil {
		panic(err)
	}

	output := number * 2

	fmt.Println(output)
}
