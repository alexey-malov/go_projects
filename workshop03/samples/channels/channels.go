package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Что make еще создает?
	c := make(chan string)

	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		// Что такое "<-" в этом случае?
		fmt.Printf("You say: %q\n", <-c)
	}

	fmt.Println("You're boring; I'm leaving.")
	close(c)
	//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}

// <- канал, в который можно только писать
func boring(msg string, c chan<- string) {
	for i := 0; ; i++ {
		// Что такое "<-" в этом случае?
		c <- fmt.Sprintf("%s %d", msg, i)

		//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
