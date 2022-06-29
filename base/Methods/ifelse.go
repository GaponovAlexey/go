package main

import (
	"fmt"
)

func main() {
	a, b := enterTheClub(222)
	fmt.Println(a, b)

}

func enterTheClub(age int) (string, bool) {
	if age >= 18 && age < 20 {
		return "Рады вас видеть", true
	} else if age >= 21 && age < 65 {
		return "Весь бар для вас", true
	} else if age >= 65 {
		return "Большой возраст для вас", false
	}
	return "Вы не проходите по фс", false
}
