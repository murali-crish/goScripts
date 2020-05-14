package main

import (
	"fmt"
	"sync"
)

func printEven(x int, wg *sync.WaitGroup) {
	if x % 2 == 0 {
		fmt.Printf("%d is even!\n", x)
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	val := 0
	for i:=0; i<10000; i++ {
		wg.Add(1)
		// go printEven(i, &wg)
		go increment(&val, &wg)
	}

	wg.Wait()

	fmt.Printf("Final value was %d\n", val)
}

func increment(ptr *int, wg *sync.WaitGroup) {
	i := *ptr
	fmt.Printf("i is %d\n", i)
	*ptr = i + 1
	wg.Done()
}