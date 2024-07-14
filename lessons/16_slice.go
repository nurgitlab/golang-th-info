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

	fmt.Println((*[3]int)(arr))
	arr2 := copy([]int{}, arr)
	arr[0] = 4

	fmt.Println(arr2)
}
func a(value []int) {
	for i, v := range value {
		value[i] = v + 2
	}
}
