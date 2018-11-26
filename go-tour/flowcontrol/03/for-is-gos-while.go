package main

import "fmt"

func main() {
	sum := 1

	// semicolons can be omitted.
	for sum < 1000 { // An equivalent to 'for ;sum < 1000; {...}'
		sum += sum
	}
	fmt.Println(sum)
}
