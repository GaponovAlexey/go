package main

import (
	"fmt"
	"time"
)

func main() {
	// append
	s := []int{}
	printSlice(s)
	s = append(s, 1, 2, 5, 6)
	printSlice(s)
	s = s[1:3]
	printSlice(s)
	//range
	FRange(data)
	fmt.Println("----")
	ContinuedRange()
	fmt.Println("----")
	m := make(map[string]Ver)
	m["Bell labs"] = Ver{40.23123123, -74.23131}
	m["Google"] = Ver{
		37.22, -44.2231,
	}
	fmt.Println(m)
	m2 := map[string]Ver{
		"test1": {22.3234, -22.445223},
		"test2": {21.3234, -21.445223},
	}
	fmt.Println(m2)
	fmt.Println("----")

	m12 := make(map[string]int)
	m12["Answer"] = 42
	fmt.Println("The value:", m12["Answer"])
	delete(m12, "Answer")
	fmt.Println("The value:", m12["Answer"])
	v, ok := m12["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	fmt.Println("----")
	t1 := map[string]int{"Go!": 1, "I": 2, "am": 1, "learning": 1}
	t2 := map[string]int{"Go!": 1, "I": 2, "am": 1, "learning": 1}
	fmt.Println(t1)
	fmt.Printf("%v\n", t2)
	//closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	//
	rr1 := time.Now().UnixNano() 
	fmt.Println(rr1)
	
}
//closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


//type
type outer struct {
	inner struct{ i int }
}

//maps
type Ver struct{ a, b float64 }

//data
var data = []int{1, 2, 3, 4, 55, 66, 77}

//append
func printSlice(s []int) {
	fmt.Printf("len%d cap=%d %v\n", len(s), cap(s), s)
}

//range
func FRange(obj []int) {
	for i, e := range obj {
		fmt.Printf("i=%d e=%d\n", i, e)
	}
}

func ContinuedRange() {
	pow := make([]int, 4)
	for i := range pow {
		pow[int(i)] = int(1) << int(i)
	}
	fmt.Println(pow)
}
