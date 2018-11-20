package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func generateArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func sortArray(arr []int, result chan<- []int) {
	sort.Ints(arr)
	result <- arr
	close(result)
}

func asyncSortArray(arr []int) <-chan []int {
	result := make(chan []int)
	go sortArray(arr, result)
	return result
}

func main() {
	arr1 := generateArray(10)
	arr2 := generateArray(5)
	arr3 := generateArray(7)

	sorted1 := asyncSortArray(arr1)
	sorted2 := asyncSortArray(arr2)
	sorted3 := asyncSortArray(arr3)

	channels := []<-chan []int{sorted1, sorted2, sorted3}
	processChannels(channels, func(index int, sorted []int) {
		fmt.Printf("arr%d%v\n", index, sorted)
	})
}

type handlerFn func(i int, arr []int)

func processChannels(chans []<-chan []int, fn handlerFn) {
	cases := make([]reflect.SelectCase, len(chans))
	for i, c := range chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(c)}
	}
	remaining := len(cases)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining--
		} else {
			fn(chosen, value.Interface().([]int))
		}
	}
}
