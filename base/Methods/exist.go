package main

import (
	"fmt"
)

func main() {
	users := map[string]int{
		"Go!":      12,
		"I":        23,
		"am":       14,
		"learning": 51,
	}
	data, ex := users["Go!"]
	if ex {
		fmt.Println(data)
	} else {
		fmt.Println("net v spiske")
	}
	fmt.Println(data, ex)
	
}
