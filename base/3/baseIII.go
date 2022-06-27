package main

import "fmt"

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
	bb:=names[1:3]
	bb[0] = "XXX"
	fmt.Println(aa, bb)
	fmt.Println(names)
	fmt.Println("----")
	
}

type Vertex struct{ X, Y int }
