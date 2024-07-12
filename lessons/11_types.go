package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

func main() {
	fmt.Println(time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
}
