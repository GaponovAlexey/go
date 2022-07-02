package main

import (
	"fmt"
)

func main() {
	user := []int{2, 4, 5, 6, 7}
	user = append(user, 33)
	fmt.Println(len(user))
	fmt.Println(cap(user))

}
