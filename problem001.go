package main
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

func main() {
	fmt.Println(euler1(1000))
}
