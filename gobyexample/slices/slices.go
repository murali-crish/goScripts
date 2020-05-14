package main

import "fmt"

func main() {

	// To create an empty slice with non-zero length, use builtin make.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the lenght of slice
	fmt.Println("len:", len(s))

	// append returns a slice containing one ore more new values.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slices can be copy'd
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// can't we just assign?
	d := c
	fmt.Println("assigned:", d)

	// slicing a slice
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slicing excluding
	l = s[:5]
	fmt.Println("sl2:", l)

	// slicing including
	l = s[2:]
	fmt.Println("sl3:", l)

	// short form slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// multidirectional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
