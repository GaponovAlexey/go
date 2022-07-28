package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	errGroup()

}

func errGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("one")
			i := 0
			for ; i <= 10; i++ {
				fmt.Println(i)
			}
		}
		return fmt.Errorf("счётчик")
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("two")
			for i := 0; i <= 10; i++ {
				fmt.Println(i)
			}
		}
		return fmt.Errorf("хуётчик")
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("two")
			for i := 0; i <= 10; i++ {
				fmt.Println(i)
			}
		}
		return fmt.Errorf("пётричик")
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("three")
			for i := 0; i <= 11; i++ {
				fmt.Println(i)
			}
		}
		return fmt.Errorf("пипи")
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
