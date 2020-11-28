package main

import "fmt"

func fac(n uint64) (ret uint64) {
	if n > 1 {
		ret = n * fac(n-1)
	} else {
		ret = 1
	}
	return ret
}
func main() {
	fmt.Println(fac(11))
}
