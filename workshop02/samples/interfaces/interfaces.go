package main

import (
	"fmt"
	"math"
)

type Printer interface {
	Print()
}

type Cat struct {
	Name string
}

//А что будет, если убрать звездочку?
func (c *Cat) Print() {
	fmt.Println(c.Name)
}

func (c Cat) String() string {
	return "Lovely cat " + c.Name
}

type Float float64

func (f Float) Print() {
	fmt.Println(f)
}

func main() {
	var i Printer

	i = &Cat{"Kitty"}
	describe(i)
	i.Print()

	i = Float(math.Pi)
	describe(i)
	i.Print()
}

func describe(i Printer) {
	fmt.Printf("(%v, %T)\n", i, i)
}
