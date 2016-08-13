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
func sumOfPrimes(n int) (sum int64) {
	if n > 1 {
		sum = 2
		for i := 3; i <= n; i += 2 {
			if isPrime(i) {
				sum =  sum + int64(i)
			}
		}

	}
	return sum
}
func euler10(n int) int64 {
	return sumOfPrimes(n)
}

func main() {
	fmt.Println(euler10(2000000))
}
