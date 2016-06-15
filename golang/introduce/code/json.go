package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	m := map[string]string{
		"html": "<>&",
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
