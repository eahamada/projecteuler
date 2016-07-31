package main

import (
	"fmt"
)

func euler6(n int) int{
  squareSum := n*n*(n+1)*(n+1)/4
  sumSquare := (2*n+1)*(n+1)*n/6
  return squareSum - sumSquare
}

func main() {
	fmt.Println(euler6(100))
}
