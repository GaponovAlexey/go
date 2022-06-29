package main

import (
	"fmt"
)

func main() {
	var m string
	m = sayHello("you", 22)
	fmt.Println(m)

}

func sayHello(n string, age int) string {
	return fmt.Sprintf("%s, %d", n, age)
}
