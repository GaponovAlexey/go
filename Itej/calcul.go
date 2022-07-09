package main

import "fmt"

func main() {
	fmt.Println("cal")
	fmt.Println("какое число (+, -, * ,/, %)")
	
	var action string
	fmt.Scan(&action)
	var a float64
	var b float64

	fmt.Println("Первое число")
	fmt.Scan(&a)
	fmt.Println("Второе число")
	fmt.Scan(&b)

	switch {
	case action == "+":
		fmt.Println("сума чисел" + fmt.Sprintln(a+b))
	case action == "-":
		fmt.Println("сума чисел" + fmt.Sprintln(a-b))
	case action == "*":
		fmt.Println("сума чисел" + fmt.Sprintln(a*b))
	case action == "/":
		fmt.Println("сума чисел" + fmt.Sprintln(a/b))
	case action == "/":
		fmt.Println("сума чисел" + fmt.Sprintln(a/b))

	}

}
