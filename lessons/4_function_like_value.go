package main

import "fmt"

func main() {
	// func можно положить в переменную
	sum := func(x, y int) int {
		return x + y
	}

	fmt.Println(sum(4, 5))
	fmt.Println(combiner(4, 5, func(x, y int) int {
		return x * y
	}))
}

func combiner(x int, y int, f func(x, y int) int) int {
	return f(x, y)
}
