package main

import "fmt"

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
