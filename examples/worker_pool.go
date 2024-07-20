package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //можно вызвать где нибудь по условию

	wg := &sync.WaitGroup{}

	to, after := make(chan int, 5), make(chan int, 5)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, to, after) //запустили воркеры
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			to <- i //пихаем данные в канал
		}
		close(to) //закрываем канал
	}()

	go func() {
		wg.Wait() //ожидаем когда выполнится WG и закрываем канал
		close(after)
	}()

	counter := 0

	for v := range after { // итеррируемся по каналам
		counter++
		fmt.Println(v)
	}

	fmt.Println("Total: ", counter)
}

func worker(ctx context.Context, to <-chan int, after chan<- int) {
	for {
		select {
		case <-ctx.Done(): // для завершения
			return
		case value, ok := <-to: //проверяем, закрыт ли канал
			if !ok {
				return
			}

			time.Sleep(time.Millisecond)
			after <- value * value
		}
	}
}
