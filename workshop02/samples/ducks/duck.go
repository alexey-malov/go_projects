package main

import "fmt"


type IDuck interface {
	Quack()
	Fly()
}
type FlyBehavior interface {
	Fly()
}

type QuackBehavior func()

type Duck struct {
	fly FlyBehavior
	quack QuackBehavior
}
type MallardDuck struct {
	Duck
}
type RubberDuck struct {
	Duck
}

func Quack(){
	fmt.Println("Quack")
}

func Squeak(){
	fmt.Println("Squeak")
}


func NewRubberDuck() *RubberDuck{
	return &RubberDuck{Duck{FlyNoWay{}, Squeak}}
}

func NewMallardDuck() *MallardDuck{
	return &MallardDuck{Duck{FlyWithWings{}, Quack}}
}

type FlyNoWay struct{}
func (FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}
type FlyWithWings struct{}
func (FlyWithWings) Fly() {
	fmt.Println("Yahooo, I am flying")
}

func (Duck) Quack() {
	fmt.Println("Quack")
}

func (d Duck) Fly() {
	d.fly.Fly()
}

/*
func (RubberDuck) Fly() {
	fmt.Println("I can't fly")
}
*/

func (RubberDuck) Quack() {
	fmt.Println("Squeak")
}


func playWithDuck(d IDuck) {
	fmt.Printf("%T\n", d)
	d.Quack()
	d.Fly()
}


func main(){

	rd := NewRubberDuck()
	md := NewMallardDuck()


	playWithDuck(md)
	playWithDuck(rd)
}

