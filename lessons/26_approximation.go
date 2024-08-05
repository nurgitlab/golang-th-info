package main

import "fmt"

type MyInt int64

func (cu MyInt) isPositive() bool {
	return cu > 0
}

type MyNumber interface {
	~int64 | float64
}

func mySum[T MyNumber](arr []T) T {
	var s T
	for _, v := range arr {
		s += v
	}

	return s
}

func main() {
	fmt.Println(mySum([]int64{1, 2, 3, 4, 5}))
	fmt.Println(mySum([]MyInt{5, 5, 5, 5, 5, 5}))
}
