package main

import "fmt"

func getMinItem(arr []int) int {
	n := len(arr)

	minItem := arr[0]

	for i := 0; i < n; i += 1 {
		if arr[i] < minItem {
			minItem = arr[i]
		}
	}
	return minItem
}

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(getMinItem(x))
}
