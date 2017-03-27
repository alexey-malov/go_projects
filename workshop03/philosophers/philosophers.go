package main

import (
	"fmt"
	"time"
)

type fork chan string

type philosopher struct {
	name        string
	lastEatTime time.Time
	leftFork    fork
	rightFork   fork
}

func newPhilosopher(name string, leftFork, rightFork fork) *philosopher {
	return &philosopher{name, time.Now(), leftFork, rightFork}
}

func start(num int) {

	forks := make([]fork, num)
	philosophers := make([]*philosopher, num)
	for i, _ := range philosophers {
		philosophers[i] = newPhilosopher(fmt.Sprint("Philosopher", i+1), forks[i], forks[(i+1)%num])
	}
	finish := make(chan string)
	for _, p := range philosophers {
		p.run(finish)
	}
	fmt.Println("Finishing simulation", <-finish, "has died from starvation")
}

func (p philosopher) think() {
	time.Sleep(time.Second * 5)
}

func (p philosopher) eat() bool {
	fmt.Println(p.name, "wants to take forks")
	p.leftFork <- p.name
	if time.Since(p.lastEatTime) > time.Second*10 {
		return false
	}
	p.rightFork <- p.name
	if time.Since(p.lastEatTime) > time.Second*10 {
		return false
	}

	fmt.Println(p.name, "is eating")

	time.Sleep(time.Second * 2)

	p.lastEatTime = time.Now()
	fmt.Println(p.name, "has finished to eat")

	if <-p.leftFork != p.name {
		panic(fmt.Sprintln("Somebody other has taken ", p.name, "left fork"))
	}
	if <-p.rightFork != p.name {
		panic(fmt.Sprintln("Somebody other has taken ", p.name, "right fork"))
	}
	return true
}

func (p philosopher) run(finish chan string) {
	go func() {
		for {
			if !p.eat() {
				finish <- p.name
				return
			}
			p.think()
		}
	}()
}

func main() {
	start(5)
}
