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

	people := []Person{
		{
			name:     "John",
			age:      1,
			jobTitle: "John",
		},
		{
			name:     "Brown",
			age:      2,
			jobTitle: "Brown",
		},
		{
			name:     "Lucky",
			age:      3,
			jobTitle: "Lucky",
		},
	}
	fmt.Println("struct:", cont(people, Person{
		name: "John",
		age:  1,
		jobTitle: "John",
	}))

}

func cont[T comparable](el []T, sel T) bool {
	for _, el := range el {
		if sel == el {
			return true
		}
	}
	return false
}
