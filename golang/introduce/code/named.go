package main

import "fmt"

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	fmt.Println(1 + 3)
}

func test() (err error) {
	defer func() {
		if err != nil {
			log.Infof("err: %v", err)
		}
	}()

	// code
}
