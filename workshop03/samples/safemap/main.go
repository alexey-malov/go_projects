package main

import (
	"fmt"
	"github.com/alexey-malov/go_projects/workshop03/samples/safemap/safemap"
	"sync"
)

func main() {
	sm := safemap.New()

	var waiter sync.WaitGroup
	const MAX_WRITERS = 1000
	waiter.Add(MAX_WRITERS)
	for i := 0; i < MAX_WRITERS; i++ {
		go func(index int) {
			key := fmt.Sprint(index)
			sm.Insert(key, "writer:"+key)
			waiter.Done()
		}(i)
	}

	waiter.Wait()
	value := sm.Close()
	fmt.Println("length:", len(value))
}
