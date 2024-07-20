package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := 0
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()

			mu.Lock()
			c++
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println(c)
}
