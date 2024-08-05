// TODO разобраться позже, похоже на промисы из JS
package main

//import (
//	"context"
//	"golang.org/x/sync/errgroup"
//)
//
////func makeRequest(num int) <-chan string {
////	res := make(chan string)
////
////	go func() {
////		time.Sleep(time.Second)
////		res <- fmt.Sprintf("Response N%d\n", num)
////	}()
////
////	return res
////}
////func main() {
////	firstRequest := makeRequest(1)
////	secondRequest := makeRequest(2)
////
////	fmt.Println(<-firstRequest, <-secondRequest)
////}
//
//func main() {
//	g, ctx := errgroup.WithContext(context.Background())
//}
