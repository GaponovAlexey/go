package main

import (
	"fmt"
)

type Users struct {
	name     string
	age      int64
	password string
}

func main() {
	user := Users{name: "Ji", age: 22, password: "23312"}
	fmt.Println(user)
	
}
