package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	buf int = runtime.NumCPU()
)

func main() {
	gracefulShutdown()
}

func gracefulShutdown() {
	sugChan := make(chan os.Signal, 1)
	signal.Notify(sugChan, syscall.SIGINT, syscall.SIGTERM)
	timer := time.After(5 * time.Second)

	for {
		select {
		case <-timer:
			fmt.Println("tim's up")

			return
		case stg := <-sugChan:
			fmt.Printf("Got message %v\n", stg)
			return
		}
	}
}
