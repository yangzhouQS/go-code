package main

import (
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	Message string `json:"message"`
}
type UserRequest struct {
	User string `json:"user"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var request UserRequest
	// 解码
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "bad err", http.StatusBadRequest)
		return
	}
	response := UserResponse{Message: "hello: " + request.User}
	// json编码
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
