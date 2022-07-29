package main

import "fmt"

type Number interface {
	int64 | float64
}
type CustInt int64

func (ci CustInt) IsPositive() bool {
	return ci > 0
}

type Numbers[T Number] []T

func main() {
	fl := []float64{1.2, 3.2, 3.6, 3.3}
	in := []int64{2, 1, 3}
	fmt.Println(sum(fl))
	fmt.Println(sum(in))

}

func sum[V int64 | float64](n []V) V {
	var sum V
	for _, v := range n {
		sum += v
	}
	return sum
}
