package main

import (
	"fmt"
	"github.com/alexey-malov/go_projects/anagrams/anagrams"
)

func main() {
	fmt.Println(anagrams.IsAnagram("ab", "ba"))
}
