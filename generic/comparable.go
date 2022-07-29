package main

import "fmt"

type Number interface {
	int64 | float64
}
type CustInt int64

func (ci CustInt) IsPositive() bool {
	return ci > 0
}

type Numbers[T Number] []T

func main() {
	showContains()

}

func showContains() {
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}
	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("int:", cont(ints, 4))

	// string := []string{"Pero", "katya", "katya", "_"}
	// fmt.Println("string:", cont(string, "katya"))
	// fmt.Println("string:", cont(string, "Pero"))

}
func cont[T comparable](el []T, sel T) bool {
	for _, el := range el {
		if sel == el {
			return true
		}
	}
	return false
}
