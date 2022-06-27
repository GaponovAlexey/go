package main

import "fmt"

func main() {
	// append
	s := []int{}
	printSlice(s)
	s = append(s, 1, 2, 5, 6)
	printSlice(s)
	s = s[1:3]
	printSlice(s)
	//range
	for i, v:= range pow {
		
	}
}
var pow = []int{1,2,3,4,55, 66, 77}
//append
func printSlice(s []int) {
	fmt.Printf("len%d cap=%d %v\n", len(s), cap(s), s)
}
//range
