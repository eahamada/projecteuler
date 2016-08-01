package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n < 9 {
		return true
	}
	if n%3 == 0 {
		return false
	}
	for f := 5; f*f <= n; f += 6 {
		if n%f == 0 {
			return false
		}
		if n%(f+2) == 0 {
			return false
		}

	}
	return true
}

func nthPrime(n int) int {
	if n < 1 {
		return -1
	}
	count := 1
	result := 2
	for i := 3; count < n; i += 2 {
		if isPrime(i) {
			result = i
			count++
		}
		
	}

	return result
}
func euler7() int {
  return nthPrime(10001)
}
func main() {
	fmt.Println(euler7())
}
