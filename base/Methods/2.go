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

	user1.printUser("Ivan")
}

func (u Tuser) printUser(name string) {
	u.name = name
	fmt.Println(u.age, u.name)

}
