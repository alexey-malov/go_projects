package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

// Каналы можно возвращать как любое обычное значение
// В данном случае возвращаем канал только для получения данных
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// возвращает доступный для чтения канал, в который ее сопрограмма пишет в фоне,
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() { // Запускаем горутину внутри функции
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Возвращаем вызывающей функции канал
}
