package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(double(a))
}

func double(nums []int) []int {
	a := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		a = append(a, nums[i]*2)
	}
	return a
}
