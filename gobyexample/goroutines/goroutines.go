package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// synchronous way of calling a function
	f("direct")

	// execute function concurrently
	go f("goroutine")

	// go routine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Wait for our two go routines to finish
	time.Sleep(time.Second)
	fmt.Println("done")
}
