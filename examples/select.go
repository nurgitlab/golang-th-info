package main

import (
	"fmt"
	"time"
)

func main() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "message1"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			message2 <- "message2"
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-message1:
			{
				fmt.Println(msg1)
			}

		case msg2 := <-message2:
			{
				fmt.Println(msg2)
			}
		}
	}
}
