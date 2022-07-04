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

type DumpDataBase struct {
	m map[string]string
}

func NewDumpDataBase() *DumpDataBase {
	return &DumpDataBase{
		m: make(map[string]string),
	}
}

func NewTUser(name, sex string, age, weight, height int) Tuser {
	return Tuser{
		name: name, sex: sex, age: age, weight: weight, height: height,
	}
}

func main() {
	user1 := NewTUser("Vasia", "male", 22, 22, 442)
	// user := Tuser{"Vasiliy", 22, "male", 76, 185}
	// user2 := Tuser{"Vasiliy", 33, "male", 76, 185}

	// fmt.Println(NewTUser())
	fmt.Printf("%+v\n", user1)

}
