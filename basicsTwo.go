package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	//for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// for init post
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	// pow
	fmt.Println(pow(9, 3, 40))
	//switch
	fmt.Println("go runs on")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Os X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%v.", os)
	}
	//day
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	//time
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	//continue last-in-first-out order
	fmt.Println("continue")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}

//pow степерь
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
