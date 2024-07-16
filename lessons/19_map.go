package main

import "fmt"

func main() {
	m := map[int]int{}

	m[0] = 10
	m[1] = 11

	fmt.Println(m)
	delete(m, 0)
	fmt.Println(m)
}
