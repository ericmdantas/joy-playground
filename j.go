package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(3 * time.Second)
	
	println("yo!")

	for {
		select {
			case <- ticker.C:
				println("hey apple!")
		}
	}
}