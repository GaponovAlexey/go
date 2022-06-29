package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := enterTheClub(222)
	fmt.Println(a, b)

}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 20 {
		return "Рады вас видеть", nil
	} else if age >= 21 && age < 65 {
		return "Весь бар для вас", nil
	} else if age >= 65 {
		return "Большой возраст для вас", errors.New("you are too old")
	}
	return "Вы не проходите по фс", errors.New("you are too child")
}
