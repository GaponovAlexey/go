package main

import (
	"fmt"
	"strings"
)

func main() {
	// pointer
	i, j := 42, 2801
	p := &i
	fmt.Println(i, j, *p)
	*p = 21
	fmt.Println(i)
	//structure
	fmt.Println(Vertex{1, 2})
	//structure
	v := Vertex{3, 3}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
	// pointers
	pp := &v
	pp.X = 1e3
	fmt.Println(v)
	//structure literals
	var (
		v1  = Vertex{1, 2}
		ppp = &Vertex{1, 2}
		v2  = Vertex{X: 1}
		v3  = Vertex{}
	)
	fmt.Println(v1, ppp, v2, v3)
	fmt.Println("----")
	//array
	var a [2]string
	a[0] = "hello"
	a[1] = "go"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	//array
	primes := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(primes)
	//slice
	var s []int = primes[0:3]
	fmt.Println(s)
	//slice pointers
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	aa := names[0:2]
	fmt.Println(aa)
	bb := names[1:3]
	bb[0] = "XXX"
	fmt.Println(aa, bb)
	fmt.Println(names)
	fmt.Println("----")
	//slice literals
	q := []float64{2, 3, 5, 67, 13}
	fmt.Println(q)
	r := []bool{true, true, true, false, true}
	fmt.Println(r)
	ss := []struct {
		i int
		b bool
	}{
		{2, true}, {3, true}, {5, true}}
	fmt.Println(ss)
	tt := []struct{ i, id int }{
		{2, 4},
		{3, 6},
	}
	fmt.Println(tt)
	//slice bounds
	sq := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(sq)
	sq = sq[3:]
	PrintSlice(sq)
	//slice nil
	var sw []int
	fmt.Println(sw, len(sw), cap(sw))
	if sw == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	//making slice
	a1 := make([]int, 4)
	printSlice("a", a1)
	// slice-of-slice string
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"
	board[2][0] = "-"
	fmt.Println(board)
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// append
	var 
}

//make-PrintSlice
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)

}

// slice-len-cap
func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

type Vertex struct{ X, Y int }
