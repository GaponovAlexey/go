package main

import "fmt"

func main() {
	n := 8

	for {
		defer fmt.Println(n)
		n = n + 1
		if n >= 12 {
			break
		}
	}
}
