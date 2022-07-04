package main

import (
	"fmt"
)

type Tuser struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func main() {
	user1 := Tuser{"Vasia", 22, "male", 76, 122}
	
	user1.PrintUser()
}

func (u Tuser) PrintUser() {
	fmt.Println(u.age)
}
