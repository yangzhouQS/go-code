package main

import (
	"encoding/json"
	"fmt"
)

type stu struct {
	Id   int
	Name string
}

func main() {
	s := stu{Id: 13, Name: "李四"}
	ret, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("%#v \n", string(ret))
	}
	fmt.Println(string(ret))
}
