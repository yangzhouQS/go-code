package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello word", r.URL.Path)
}
func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
