package main

import (
	"fmt"
)

func largestPrimeFactor(n, p int64) int64 {
	if p*p > n {
		return n
	}
	var _n = n
	var _p = p
	if n%p == 0 {
		_n = n / p
		for _n%p == 0 {
			_n = n / p
		}
	}

	if p > 2 {
		_p += 2
	} else {
		_p += 1
	}
	return largestPrimeFactor(_n, _p)
}

func euler3(n int64) int64 {
	return largestPrimeFactor(n, 2)
}

func main() {
	fmt.Println(euler3(600851475143))
}
