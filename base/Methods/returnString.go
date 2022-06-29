package main

import (
	"fmt"
)

func main() {
	fmt.Println(sayHello("you", 22))

}

func sayHello(n string, age int) string {
	return fmt.Sprintf("%s, %d", n, age)
}
