package main

import (
	"fmt"
	mySum "projects/go-road/lessons/27_packages/sum"
)

func main() {
	fmt.Println("Sum")

	fmt.Println(mySum.SumInt(1, 2, 3, 4))
}
