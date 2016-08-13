package main

import (
	"fmt"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func euler9(s int) (a,b,c int){

	s2 := s / 2
	for m := 2; m*m <= s2; m++ {
		if s2%m == 0 {
			sm := s2 / m
			for sm%2 == 0 {
				sm /= 2
			}
			k := m + 1
			if m%2 == 1 {
				k += 1
			}
			for k < 2*m && k <= sm {
				if sm%k == 0 && gcd(k, m) == 1 {
					d := s2 / (k * m)
					n := k - m
					a = d * (m*m - n*n)
					b = 2 * d * m * n
					c = d * (m*m + n*n)

				}
				k += 2
			}
		}

	}
    return a, b, c
}

func main() {
	fmt.Println(euler9(1001))
}
