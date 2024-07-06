package main

import "fmt"

func main() {
	m := make(chan string, 3)
	m <- "hello1"
	m <- "hello2"
	m <- "hello3"
	fmt.Println(<-m)

	m <- "hello4"

	fmt.Println(<-m)
}

//func main() {
//	message := make(chan string)
//	go createNums(message)
//
//	for {
//		msg, isOpen := <-message
//		if !isOpen {
//			break
//		}
//
//		fmt.Printf(msg)
//	}
//}
//
//func createNums(m chan string) {
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Second)
//		m <- fmt.Sprintf("%d", i)
//	}
//
//	close(m)
//}
