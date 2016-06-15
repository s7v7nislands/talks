package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var count = 0

func Count() {
	mu.Lock()
	defer mu.Unlock()

	count += 1
}

func main() {
	for i := 0; i < 1000; i++ {
		go Count()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(count)
}
