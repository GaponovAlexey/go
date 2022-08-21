package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]interface{}{}
	m["one"] = 1
	m["two"] = 2.0
	m["three"] = true
	for k, v := range m {
		switch v.(type) {
		case bool:
			fmt.Printf("%v is an bool %v\n", k, v)
		case int:
			fmt.Printf("%v is an int %v\n", k, v)
		case float64:
			fmt.Printf("%v is an float64 %v\n", k, v)
			default:
			fmt.Printf("%v os %v\n", k, reflect.TypeOf(v))

		}
	}
}
