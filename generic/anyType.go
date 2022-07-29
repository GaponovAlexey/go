package main

import "fmt"


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
	strin := []string{"Pero", "katya", "katya", "_"}

	fmt.Println(ints)
	show(ints)
	show(strin)
	show("hello", "katya")
	show([]int64{1, 2}, []int64{1, 2, 3})
	show([]string{"A", "B", "C"}, []string{"A", "B", "C"})
	show(map[string]int64{
		"a": 1,
		"b": 2,
	})
	show(interface{}(1), interface{}("string"), any(struct{name string}{name: "Ji"}))
}



func show[T any](el ...T) {
	fmt.Println(el)

}
