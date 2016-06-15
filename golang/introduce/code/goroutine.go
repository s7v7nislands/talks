package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go hello(i, done)
	}

	time.Sleep(3 * time.Second)
	close(done)
	fmt.Println("Bye!")
}

func hello(num int, done <-chan struct{}) {
	fmt.Printf("Hello World! %d\n", num)
	<-done
}
