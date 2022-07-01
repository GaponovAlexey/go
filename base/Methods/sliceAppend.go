package main

import (
	"fmt"
)

func main() {
	mes := make([]string, 5) // cap 5
	mes = append(mes, "5")   // cap 10
	mes = append(mes, "6")
	mes = append(mes, "7")
	mes = append(mes, "8")
	mes = append(mes, "9")
	mes = append(mes, "10") //cap 20
	// cap() - Огранчение 10000 пл 10001
	fmt.Println(len(mes))
	fmt.Println(cap(mes))
	fmt.Println(mes)

}
