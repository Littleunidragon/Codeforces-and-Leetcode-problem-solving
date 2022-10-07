package main

import (
	"fmt"
	"strconv"
)

/*func EmojiArray(n int) []string {
	ans := make([]string, n)
	for i := range ans {
		ans[i] = string(i)
	}
	return ans
}*/

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 0; i != n; i++ {
		j := i + 1
		if j%3 == 0 && j%5 != 0 {
			ans[i] = "Fizz"
		} else if j%5 == 0 && j%3 != 0 {
			ans[i] = "Buzz"
		} else if j%3 == 0 && j%5 == 0 {
			ans[i] = "FizzBuzz"
		} else {
			ans[i] = fmt.Sprint(j)
		}
	}
	return ans
}

func main() {
	fmt.Println(fizzBuzz(5))
}

func BetterfizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 0; i < n; i++ {
		j := i + 1
		switch true {
		case j%3 == 0 && j%5 == 0:
			answer[i] = "FizzBuzz"
		case j%3 == 0:
			answer[i] = "Fizz"
		case j%5 == 0:
			answer[i] = "Buzz"
		default:
			answer[i] = strconv.Itoa(j)
		}
	}
	return answer

}
