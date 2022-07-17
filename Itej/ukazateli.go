package main

import (
	"fmt"
)



func main() {
	a:= 5
	
	p:= &a
	*p = 2
	fmt.Println(*p)
	fmt.Println(a)
	fmt.Println(a)
	
}
