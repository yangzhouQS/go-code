package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserResponse struct {
	Message string `json:"message"`
}
type UserRequest struct {
	User string `json:"user"`
}

func HelloHandlerRes(w http.ResponseWriter, r *http.Request) {
	// body体参数
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "res fail", http.StatusBadRequest)
		return
	}
	// 请求数据
	var req UserRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "res fail", http.StatusBadRequest)
		return
	}
	// 响应数据
	res := UserResponse{Message: "世界你好: " + req.User}
	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}
func main() {
	http.HandleFunc("/hello", HelloHandlerRes)
	http.ListenAndServe(":8080", nil)
}
