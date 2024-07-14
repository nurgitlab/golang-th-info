package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 10
	b := 3

	c := float64(a)
	d := float64(b)

	fmt.Println(c / d)

	fmt.Println("_____")

	fmt.Println(strconv.Itoa(a) + "String!") //to string

	fmt.Println(strconv.Atoi("fff"))

}
