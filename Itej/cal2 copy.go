package main

import "fmt"

type Users struct {
	name  string
	age   float64
	score []int
}

func(u Users) getHi() int {
	hight:= 0

	for _, sc := ranger u.score {
		
	}
}

func main() {
	data := Users{"ivan", 22, []int{1, 2, 3}}
	fmt.Println(data.score)

}
