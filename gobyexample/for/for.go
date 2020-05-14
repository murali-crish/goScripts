package main

import "fmt"

func main() {

	// The most basic type, with single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// infinite loop until break out of loop or return from enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
