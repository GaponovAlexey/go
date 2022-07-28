package main

import "fmt"

func main() {
	num := []int{5, 5, 6, 7, 5, 11}
	c := make(chan int)

	go func() {
		for _, i := range num {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println("chan:", v)

	}
}
