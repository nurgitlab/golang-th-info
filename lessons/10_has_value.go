package main

import "fmt"

func main() {
	m := map[int]int{}

	m[1] = 1

	if a, isHas := m[2]; isHas {
		fmt.Println(a)
	} else {
		fmt.Println("Not found")
	}
}
