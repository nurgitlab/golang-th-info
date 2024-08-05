package main

import "fmt"

func main() {
	//fmt.Println(showSum(5, 3))
	//fmt.Println(showSum(6.0, 7.0))

	//fmt.Println(isHas([]int{1, 2, 4, 5}, 4))
	//fmt.Println(isHas([]int{1, 2, 4, 5}, 44))
	//fmt.Println(isHas([]string{"r", "rr"}, "rrr"))

	fmt.Println(sum(1, 2, 4, 5))
}

type Number interface {
	int | int64 | float32 | float64
}

func sum[V Number](nums ...V) V {
	var sum V = 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

//func isHas[T comparable](arr []T, target T) bool {
//	for _, item := range arr {
//		if item == target {
//			return true
//		}
//	}
//
//	return false
//}

//func showSum[T int | float64](a, b T) T {
//	return a + b
//}
