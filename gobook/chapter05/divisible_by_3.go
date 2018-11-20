package main

import "fmt"

func main() {
	for i := 1; i <= 100; i += 1 {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
