package main
import "fmt"

func fib(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	return fib(n-1) + fib(n-2)
}

func _euler2(max, n, acc int) int {
	if fib(n) > max {
		return acc
	}
	fib := fib(n)
	if fib%2 == 0 {
		acc += fib
	}
	return _euler2(max, n+1, acc)
}

func euler2(max int) int {
	return _euler2(max, 0, 0)
}

func main() {
	fmt.Println(euler2(4000000))
}
