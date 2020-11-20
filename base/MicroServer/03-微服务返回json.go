package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"user"`
}

func HelloHanderX(w http.ResponseWriter, r *http.Request) {
	user := User{31, "李四"}
	//marshal, _ := json.Marshal(user)
	//fmt.Fprint(w, string(marshal))
	// 使用json编码
	decoder := json.NewEncoder(w)
	decoder.Encode(user)
}
func main() {
	port := 8080
	http.HandleFunc("/hello", HelloHanderX)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
