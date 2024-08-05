package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//mutex()
	addAtomic()
}

// атомик работает на уровне процессора
func addAtomic() {
	counter := int64(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		atomic.AddInt64(&counter, 1)
	}

	fmt.Println("Counter:", counter)
}

func mutex() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	c := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			c++
			mu.Unlock()
			defer wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(c)
}
