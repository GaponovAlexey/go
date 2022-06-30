package main

import "fmt"

func main() {
	x:= 0
	fmt.Println(&x)
	z:= &x
	fmt.Println(*z)
	
}
