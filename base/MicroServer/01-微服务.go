package main

import (
	"fmt"
	"net/http"
)

func HelloWordHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello word!")
}
func main() {
	port := 8080
	http.HandleFunc("/hello", HelloWordHander)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
