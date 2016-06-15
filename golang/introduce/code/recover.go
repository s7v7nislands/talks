package main

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("in defer")
			log.Printf("recover: %v\n", err)
		}
	}()

	die()

}

func die() {
	panic("died")
}
