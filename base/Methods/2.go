package main

import "fmt"

func main() {
	fmt.Println(enterTheClub("mo"))

}

func enterTheClub(day string) string {
	switch day {
	case "mo":
		return "monday"
	case "tu":
		return "tuesday"
	case "we":
		return "wednesday"
	default:
		return "you"
	}
}
