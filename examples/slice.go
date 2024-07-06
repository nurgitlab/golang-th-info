package main

import "fmt"

func main() {
	var list []int
	fmt.Println(list == nil) //true

	list = []int{}

	fmt.Println(list == nil) //false
}
