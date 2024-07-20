package main

import (
	"context"
	"fmt"
	"time"
)

// WORKER POOL
func main() {
	c := context.Background() //корневой контекст
	d := context.TODO()       //для тестов, не нужен
	fmt.Println(c)
	fmt.Println(d)

	//withValue := context.WithValue(c, "key", "value")
	//fmt.Println(withValue)
	//fmt.Println(withValue.Value("key"))

	//withCancel, cancel := context.WithCancel(c)
	//fmt.Println(withCancel.Err())
	//cancel()
	//fmt.Println(withCancel.Err())

	//withDeadline, cancel := context.WithDeadline(c, time.Now().Add(time.Second*5))
	//defer cancel()
	//fmt.Println(withDeadline.Deadline())
	//fmt.Println(withDeadline.Err())
	//fmt.Println(<-withDeadline.Done())
	//fmt.Println("Here")

	withTimeout, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()
	fmt.Println(withTimeout.Done())
}
