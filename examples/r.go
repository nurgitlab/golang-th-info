package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	var wg = sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w(ctx, in, out)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	counter := 0
	for v := range out {
		fmt.Println(v)
		counter++
	}
	fmt.Println("counter", counter)
}

func w(ctx context.Context, in <-chan int, out chan<- int) {
	for {
		select {
		case v, ok := <-in:
			if ok {
				time.Sleep(time.Millisecond)
				out <- v + 10
			} else {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
