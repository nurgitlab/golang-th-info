package main

import (
	"fmt"
	"time"
)

func main() {
	resultChan := make(chan int)
	timer := time.After(time.Second)

	go func() {
		defer close(resultChan)
		for i := 0; i < 1000; i++ {
			select {
			case <-timer:
				fmt.Println("timer up")
				return
			default:
				time.Sleep(1 * time.Millisecond)
				resultChan <- i
			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}
