package main

import (
   "fmt"
   "net/http"

)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go is super easy!")}

func main() {
	fmt.Println("start: 3000") 
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":3000", nil)
}