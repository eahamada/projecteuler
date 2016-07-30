// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func max(a,b int) int{
   result :=a
   if(a < b){
     result =b
   }
   return result;
}

func expBySquaring(x, n int) int {
	return _expBySquaring(1, x, n)
}

func _expBySquaring(y, x, n int) int {
	if n < 0 {
		return _expBySquaring(y, 1/x, -n)
	} else if n == 0 {
		return y
	} else if n == 1 {
		return x * y
	} else if n%2 == 0 {
		return _expBySquaring(y, x*x, n/2)
	} else {
		return _expBySquaring(x*y, x*x, (n-1)/2)
	}
}

func isPalindromicString(s string) bool {
	length := len(s)
	if length == 1 {
		return true
	}

	if s[0] == s[length-1] {
		if length <= 3 {
			return true
		} else {
			return isPalindromicString(s[1 : length-1])
		}
	}
	return false
}

func isPalindromicInt(n int) bool {
	if n < 0 {
		return isPalindromicInt(-n)
	}
	if n < 10 {
		return true
	}
	return isPalindromicString(strconv.Itoa(n))
}

func largestPalindromicNumberFromProductOfNDigits(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be strictly positive")
	}
	result := 9
	if n == 1 {
		return 9, nil
	} else {
		for i := expBySquaring(10, n-1) ; i < expBySquaring(10, n); i++ {
			for j := expBySquaring(10, n-1) ; j < expBySquaring(10, n); j++ {
			  product := i*j
			  if(isPalindromicInt(product)){
				result = max(result,product);
			  }
			}
		}
	}
	return result, nil
}

func euler4(n int) (int, error) {
	return largestPalindromicNumberFromProductOfNDigits(n)
}

func main() {
	fmt.Println(euler4(3))
}
