package main

import "fmt"

func main() {

	fmt.Println("1") //в обратном порядке
	myPanic()
	fmt.Println("2")
	//fmt.Println("Количество потоков:", runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	////параллельно до 10 горутин
	//
	//for i := 0; i < 10; i++ {
	//	go slow(i)
	//}
	//
	//time.Sleep(15 * time.Second)
}
func myPanic() {
	defer func() {
		panicValue := recover()
		fmt.Println("->", panicValue)
	}()
	panic("something bad happened")
}

//func slow(i int) {
//	time.Sleep(1 * time.Second)
//	fmt.Printf("N:%v \n", i)
//}
