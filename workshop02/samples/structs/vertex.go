package main

import "fmt"

type Vertex struct {
	X int // поле доступно в пакетах
	Y int
}

//Илья, расскажи про указатели ;)
//И про вывод структур
//И про регистр имен полей напомни еще раз

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
