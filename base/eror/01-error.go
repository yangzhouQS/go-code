package main

import (
	"errors"
	"fmt"
)

func test(a, b int) (ret float64, err error) {
	if b == 0 {
		err = errors.New("除数不可以为0")
		return
	} else {
		ret = float64(a / b)
		return
	}
}

func main() {
	ret, err := test(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}
