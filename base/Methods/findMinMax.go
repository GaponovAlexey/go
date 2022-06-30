package main

import "fmt"

func main() {
	fmt.Println(findMin(2, 3, 4))
	fmt.Println(findMax(3, 5, 11, 22, 33))

}

func findMin(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	min := n[0]
	for _, i := range n {
		if i < min {
			min = i
		}
	}
	return min
}
func findMax(z ...int) int {
	if len(z) == 0 {
		return 0
	}
	max := z[0]
	for _, i := range z {
		if i > max {
			max = i
		}
	}
	return max

}
