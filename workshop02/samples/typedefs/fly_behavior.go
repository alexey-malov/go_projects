package main

import "fmt"

type myInt int
type MyString string
type FlyBehavior func(direction myInt) MyString

func main() {
	var noFly FlyBehavior = func(direction myInt) MyString {
		return MyString(fmt.Sprintf("no fly! %v", direction))
	}

	result := noFly(10)
	fmt.Println(result)
}
