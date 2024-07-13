package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	double(&arr)

	fmt.Println(arr)

	m := make([]int, 0, 10)

	fmt.Println(m)
}

func double(arr *[]int) {
	p := *arr

	for i := 0; i < len(p); i++ {
		p[i] *= 2
	}
}
