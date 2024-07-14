package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 3, 2, 6, 9, 2, 6, 4, 1, 5}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr)
}
