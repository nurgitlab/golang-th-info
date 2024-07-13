package main

import "fmt"

// неограниченное число элементов
//func showAllElements(values ...int) {
//	for i := 0; i < len(values); i++ {
//		fmt.Println(values[i])
//	}
//}

func main() {
	arr := []int{1, 2, 3}
	arr2 := []int{1, 2, 3}

	fmt.Println(append(arr, arr2...))
}
func a(value []int) {
	for i, v := range value {
		value[i] = v + 2
	}
}
