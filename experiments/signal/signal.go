package main

import "fmt"

type weather struct {
	temp          float64
	tempObservers []chan<- float64
}

func makeWeather() weather {
	w := weather{0.0, make([]chan<- float64, 0, 10)}
	return w
}

func (w *weather) onTemperatureChange(ch chan<- float64) {
	w.tempObservers = append(w.tempObservers, ch)
}

func (w *weather) setTemperature(t float64) {
	w.temp = t
	for _, ch := range w.tempObservers {
		ch <- t
	}
}

func main() {
	w := makeWeather()

	ch := make(chan float64)
	w.onTemperatureChange(ch)

	go func() {
		for {
			select {
			case t, ok := <-ch:
				if !ok {
					fmt.Println("Closed")
					return
				}
				fmt.Println("Temp:", t)
			}
		}
	}()

	w.setTemperature(42)
	w.setTemperature(33)
}
