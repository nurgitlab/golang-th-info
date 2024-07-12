package main

import "fmt"

func main() {
	second := first(10)

	fmt.Println(second(1))
	fmt.Println(second(2))
	fmt.Println(second(4))
}

func first(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
