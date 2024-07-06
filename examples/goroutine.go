package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//go fmt.Println("Go Go")
	//fmt.Println("Goroutine")
	//
	//time.Sleep(1 * time.Millisecond)

	t := time.Now()
	rand.Seed(t.UnixNano())
	fmt.Println(t)

	go parseUrl("111")
	parseUrl("222")

	fmt.Println("Total ", time.Since(t).Seconds())
}

func parseUrl(url string) {
	for i := 0; i < 5; i++ {
		latency := rand.Intn(500) + 500

		time.Sleep(time.Duration(latency) * time.Millisecond)

		fmt.Printf("Parsing <%s>, i= %d - Latency= %d ms\n", url, i, latency)
	}
}
