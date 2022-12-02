package main

import (
	"fmt"

)

func main() {
	i := 2
	describe(i)

	d := "tt"
	describe(d)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
