package main

import (
	"fmt"
)

type Age int

func main() {
	// user1 := Tuser{"Vasia", 22, "male", 76, 122}

	fmt.Println(Age.isAge(22))

}

func (a Age) isAge() bool {
	return a >= 20
}
