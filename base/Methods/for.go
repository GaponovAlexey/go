package main

import (
	"fmt"
)

func main() {
	matrix := make([][]int, 10)
	fmt.Println(matrix)

	conut := 0
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			conut++
			matrix[x][y] = conut
		}
		fmt.Println(matrix[x])

	}
}
