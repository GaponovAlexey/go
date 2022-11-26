package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	withoutErrGroup()
}

func withoutErrGroup() {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("second started")
			err = fmt.Errorf("any error")
			cancel()
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println(err)
}