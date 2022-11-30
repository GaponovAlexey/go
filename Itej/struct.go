package main

import (
	"fmt"
	"log"
)

type Users struct {
	name     string
	age      int64
	password string
}

func main() {
	user := Users{name: "Ji", age: 22, password: "23312"}
	fmt.Println(user)
	pilot := Users{name: "ivan", age:101, password: "123"}
	log.Println(pilot)
	
	
}
