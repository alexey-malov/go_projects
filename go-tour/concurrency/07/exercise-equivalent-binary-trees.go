package main

import (
	"fmt"
	"github.com/golang/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	walk := func(t *tree.Tree, ch chan int) {
		Walk(t, ch)
		close(ch)
	}
	go walk(t1, ch1)
	go walk(t2, ch2)

	for v1 := range ch1 {
		select {
		case v2 := <-ch2:
			if v1 != v2 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
