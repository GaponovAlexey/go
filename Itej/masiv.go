package main

import (
	"fmt"
	"math"
)

func main() {
	marks := [5]float64{1, 2, 3, 4, 5}
	var sum float64
	// for i := 0; i < len(marks); i++ {
	// 	sum += marks[i]
	// }
	for _, v := range marks {
		sum += v
	}
	fmt.Println(sum)
	var res float64 = sum / float64(len(marks))
	fmt.Println(math.Round(res))
}
