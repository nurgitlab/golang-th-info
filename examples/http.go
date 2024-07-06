package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://echo.free.beeceptor.com/",
		"google.ru",
		"yandex.ru",
	}

	//Wait Group
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func() {
			doHttp(urls[i])
			wg.Done()
		}()
	}

	wg.Wait()
}

func doHttp(url string) {
	time.Sleep(1000 * time.Millisecond)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close() //нужно закрыть поток для чтения
	fmt.Println("Status:", resp.Status)
}
