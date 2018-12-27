package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(s float64) {
	v.X *= s
	v.Y *= s
}

type Vertex3D struct {
	Vertex
	Z float64
}

func (v *Vertex3D) Scale(s float64) {
	v.Vertex.Scale(s)
	/*
		v.X *= s
		v.Y *= s
	*/
	v.Z *= s
}

type Line struct {
	start, end Vertex
}

func main() {
	v1 := Vertex3D{Vertex{3, 4}, 5}
	v2 := Vertex3D{}
	v2.X = 1
	v2.Y = 2

	fmt.Println(v1.X)
	fmt.Println(v2)

	line := Line{Vertex{1, 1}, Vertex{2, 3}}
	fmt.Println(line)
}
