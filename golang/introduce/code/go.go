package main

import "fmt"

func add(x, y int) int {
	fmt.Println(x, y)
	return x + y
}
func main() {

	add(1, 2)
	go add(1, 2)
}
