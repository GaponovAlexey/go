package main

import (
	"fmt"
)

func main() {
	a, b := enterTheClub(33)
	fmt.Println(a, b)

}

func enterTheClub(age int) (string, bool) {
	if age >= 18 && age < 21 {
		return "Рады вас видеть", true
	} else if age >= 22 {
		return "Весь бар для вас", true
	}
	return "Вы не проходите по фс", false
}
