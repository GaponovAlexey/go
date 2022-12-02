package main

import (
	"fmt"

)

type Vertex struct {
	X, Y string
}

func main() {
	b := make([]int, 5, 5) // len(b)=0, cap(b)=5
	b = append(b, 10, 11)
	b[0] = 10

	fmt.Println("B = ", b, len(b), cap(b))

	c := make([]int, len(b))
	copy(c, b)

	fmt.Println("\nC = ", c, len(c), cap(c))
	fmt.Println(b)
	fmt.Println(c)

}
