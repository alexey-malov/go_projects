package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Вызвать как горутинку
// Что произошло?

func main() {
	go boring("boring!")
	time.Sleep(3 * time.Second)
}

func boring(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
