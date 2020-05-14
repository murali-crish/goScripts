package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 0; i < 10; i++ {
		go fmt.Printf("Goroutine number: %d\n", i)
	}
	fmt.Println("Loop Finished")
	time.Sleep(1 * time.Second)
}
