package main

import (
	"fmt"

)

func goChan(i []int, c chan int) {
	sum := 0
	for _, s := range i {
		sum += s
	}
	c <- sum
}

func main() {

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go goChan(s, c)
	go goChan(s, c)
	x, f := <-c, <-c

	fmt.Println(x,f, x+f)

}
