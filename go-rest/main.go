package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2}
	s := a[:]
	s[0] = 42
	fmt.Printf("%#v\n", a) // []int{42, 2}
}
