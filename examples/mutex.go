package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type intLock struct {
	val int
	sync.Mutex
}

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	mat := [][]int{}

	//Create Matrix
	for i := 0; i < 2; i++ {
		line := []int{}
		for j := 0; j < 2000000; j++ {
			line = append(line, j)
		}
		mat = append(mat, line)
	}

	//fmt.Println(countNegatives(mat))
	//2186 1880 1663 1533 1341 1275 1442 1585

	fmt.Println(stupidCount(mat))
	//2331 2134 1706 1918 1676 1707 1808 1855

	fmt.Printf("Прошло %d миллисекунды", time.Since(t).Milliseconds())
}

func stupidCount(grid [][]int) int {
	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				ans++
			}
		}
	}

	return ans
}

func countNegatives(grid [][]int) int {
	n := &intLock{val: 0}

	var wg sync.WaitGroup

	for i := 0; i < len(grid); i++ {
		wg.Add(1)
		go func() {
			a := 0

			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] < 0 {
					a++
				}
			}

			n.Lock()
			n.val += a
			n.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return n.val
}
