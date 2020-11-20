package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

//只能被1 和 其本身整除的数
func main() {
	for i := 0; i < 10000; i++ {
		if isPrime(i) {
			fmt.Printf("%d 是素数", i)
		}
	}
}
