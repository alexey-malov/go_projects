package main

import "fmt"

func fmtBottles(i int) string {
	if i == 0 {
		return "бутылок"
	} else if i == 1 {
		return "бутылка"
	} else if i <= 4 {
		return "бутылки"
	} else if i <= 20 {
		return "бутылок"
	} else {
		return fmtBottles(i % 10)
	}

}

func main() {
	maxBottles := 500
	for i := maxBottles; i > 0; i-- {
		bottles := fmtBottles(i)
		fmt.Println(i, bottles, "пива на стене")
		fmt.Println(i, bottles, "пива!")
		fmt.Println("Возьми одну, пусти по кругу")
		fmt.Println(i-1, fmtBottles(i-1), "пива на стене!\n")
	}

	fmt.Println("Нет бутылок пива на стене!")
	fmt.Println("Нет бутылок пива!")
	fmt.Println("Пойди в магазин и купи ещё,")
	fmt.Println(maxBottles, "бутылок пива на стене!")
}
