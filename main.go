package main

import "fmt"

func main() {

	m := map[string]interface{}{}

	st := []string{"you", "lllea", "ki"}
	in := []int{1, 2, 5}

	for s, i := range m {
		if _, ok := m[s]; !ok {
			m = s, i = st
		}
	}
	fmt.Println(m)

}
