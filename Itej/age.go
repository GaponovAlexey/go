package main

import "fmt"

type Users struct {
	name string
	age  float64
}

func (u *Users) setName(name1 string) {
	u.name = name1
}

func (u Users) isElder() bool {
	a := u.age
	isTrue := false
	switch {
	case a >= 18:
		isTrue = true
	case a < 18:
		isTrue = false
	}
	return isTrue
}

func main() {
	data := Users{"ivan", 22}
	fmt.Println("сколько вам лет")
	fmt.Scan(&data.age)
	if data.isElder() {
		fmt.Println("zahodi")
	} else {
		fmt.Println("get")

	}
}
