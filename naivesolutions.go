import "fmt"

func isMult(a, b int) bool {
	return a%b == 0
}
func _euler1(n, acc int) int {
	if n == 0 {
		return acc
	} else {
		if isMult(n, 3) || isMult(n, 5) {
			acc += n
		}
		return _euler1(n-1, acc)
	}
}
func euler1(n int) int {
	return _euler1(n-1, 0)
}

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
	//	fmt.Println(euler1(1000))
	fmt.Println(euler2(4000000))
}
