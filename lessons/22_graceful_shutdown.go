package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //подписались на некие сигналы
	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("timed out")
		return
	case sig := <-sigs:
		fmt.Println("received signal", sig)
		return
	}
}
