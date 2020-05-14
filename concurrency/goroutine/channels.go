package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() { ch <- lucasoid(0, 1, 20) }()
	go func() { ch <- lucasoid(0, 1, 30) }()
	go func() { ch <- lucasoid(0, 1, 40) }()


	fmt.Printf("First finisher (#3) return %d\n", <-ch)
}

func lucasoid(a, b, n int) int {
	if n == 0 {
		return a
	}

	if n == 1 {
		return b
	}

	return lucasoid(a, b, n-1) + lucasoid(a, b, n-2)
}
