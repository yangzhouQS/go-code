package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://book.douban.com/"
	rep, err := http.Get(url)
	defer rep.Body.Close()
	if err != nil {
		panic(err)
	}
	if rep.StatusCode != http.StatusOK {
		fmt.Println("response status code :", rep.StatusCode)
	}
	ret, err2 := ioutil.ReadAll(rep.Body)
	if err2 != nil {
		fmt.Println("error code")
	}
	fmt.Println(string(ret))
}
