package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 字段名称必须大写
type HelloWordResponse struct {
	Id      int
	Message string
}

func HelloWordHanderJson(w http.ResponseWriter, r *http.Request) {
	response := HelloWordResponse{12, "Hello 日天"}
	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Add("content-type", "application/json;charset=UTF-8")
	w.Header().Add("userName", "xianjs")
	fmt.Fprint(w, string(data))
}
func main() {
	port := 8080
	http.HandleFunc("/hello", HelloWordHanderJson)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
