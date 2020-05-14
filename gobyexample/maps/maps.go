package main

import "fmt"

func main() {

	// empty map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// removes key/value pair
	delete(m, "k2")
	fmt.Println("map:", m)

	// optional second return value indicates if the key was present in the map.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	_, ok := m["k1"]
	fmt.Println("ok:", ok)

	// Can also declare and initialize a new map in same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
