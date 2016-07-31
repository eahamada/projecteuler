package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	result := false
	if n == 2 {
		result = true
	}
	if n > 2 && n%2 != 0 {
		result = true
		for j := 3; j*j <= n; j += 2 {
			if n%j == 0 {
				result = false
				break
			}
		}
	}
	return result
}
func allPrimesLTE(n int) []int {
	result := []int{}
	if n >= 2 {
		result = append(result, 2)
		for i := 3; i <= n; i += 2 {
			if isPrime(i) {
				result = append(result, i)
			}
		}
	}

	return result
}

func smallestMultiple(k int) int {
	p:=allPrimesLTE(k)
	a:=make([]int, len(p))
	for j := 0; j < len(a); j++ {
	  a[j] = 1
	}
	result := 1
	check := true
	for i := 0;  i < len(p); i++ {
		a[i] = 1
		if check {
			if p[i]*p[i] <= k {
				a[i] = int(math.Floor(math.Log(float64(k)) / math.Log(float64(p[i]))))
			} else {
				check = false
			}
		}
		result *= int(math.Pow(float64(p[i]), float64(a[i])))
	}
	return result
}

func main() {
	fmt.Println(smallestMultiple(20))
}
