package main

import "fmt"

var i, j int = 1, 2

func main() {
	// If an initializer is present, the type can be omitted
	var c, python, java = true, false, "No!"

	fmt.Println(i, j, c, python, java)
}
