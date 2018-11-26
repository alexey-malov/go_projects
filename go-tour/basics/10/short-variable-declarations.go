package main

import "fmt"

// Outside a function, every statement begins with a keyword (var, func, and so on) and so
// the := construct is not available
// This code won't compile
// bad := 42

func main() {
	var i, j int = 1, 2

	// Inside a function the short assignment := can be used in place of a var declaration with implicit syntax
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
