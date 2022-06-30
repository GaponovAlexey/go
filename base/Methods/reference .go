package main

import "fmt"

func main() {
	num := 0
	var p *int
	p = &num
	fmt.Println(p)
	fmt.Println(&num)

	msg := "ChangeMassage"
	fmt.Println(msg)
	ChangeMassage(&msg)
	fmt.Println(msg)
}

func ChangeMassage(msg *string) string {
	*msg += "(from is function)"
	return *msg
}
