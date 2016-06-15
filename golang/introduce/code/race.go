package main

import "time"

func main() {
	a := 1

	go func() {
		a++
	}()

	a++

	time.Sleep(5 * time.Second)
}
