package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//a := arr[2:5] //ссылается на исходный массив (слайс новый, ссылка на массив та же)
	//arr[3] = 11
	//a[2] = 100
	//fmt.Println(a)
	//fmt.Println(arr)

	//a := make([]int, len(arr))
	//copy(a, arr)
	//a[1] = 100
	//a = append(a, 99)
	//fmt.Println(a, arr)

	a := append(arr[:2], arr[3:]...) //ломает исходный слайс
	fmt.Println(a)

	fmt.Println(arr)
}
