package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	mat := [][]int{}

	//Create Matrix
	for i := 0; i < 20000; i++ {
		line := []int{}
		for j := 0; j < 20000; j++ {
			line = append(line, j)
		}
		mat = append(mat, line)
	}
	//1924мс

	//for i := 0; i < 20000; i++ {
	//	doubleLine(mat[i])
	//}
	////2735мс

	var wg sync.WaitGroup
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		go func() {
			doubleLine(mat[i])
			wg.Done()
		}()
	}
	wg.Wait()
	//2212мс

	fmt.Printf("Прошло %d миллисекунды", time.Since(t).Milliseconds())
}

func doubleLine(line []int) {
	for i := 0; i < len(line); i++ {
		line[i] = line[i] * 2
	}
}
