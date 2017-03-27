package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(300 * time.Millisecond)
	fmt.Println(fib(40))
}

func spinner(delay time.Duration) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}