package main

import (
	"fmt"
	"time"
)

func main() {
	buffered := make(chan int, 2)

	arr := []int{5, 6, 7, 8, 9, 10}

	go func() {
		for _, v := range arr {
			time.Sleep(time.Second)
			buffered <- v
		}
		close(buffered)
	}()

	for v := range buffered {
		fmt.Println(v)
	}
}

//func main() {
//	buffered := make(chan int, 2)
//	fmt.Printf("len: %v, cap: %v\n", len(buffered), cap(buffered))
//	buffered <- 1
//	buffered <- 2
//
//
//	fmt.Println(<-buffered)
//}

//func main() {
//	channel := make(chan int)
//
//	go func(write chan<- int) {
//		time.Sleep(time.Second)
//
//		write <- 1
//	}(channel)
//
//	value := <-channel
//	fmt.Println(value)
//
//	go func(read <-chan int) {
//		time.Sleep(time.Second)
//
//		fmt.Println(<-read)
//	}(channel)
//	fmt.Println(channel)
//	channel <- 2
//
//	channel <- 3 //deadlock!
//}
