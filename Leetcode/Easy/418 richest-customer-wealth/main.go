package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	var sum, ans int
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[0]); j++ {
			sum += accounts[i][j]
		}
		if sum > ans {
			ans = sum
		}
		sum = 0
	}
	return ans
}

func BettermaximumWealth(accounts [][]int) int {
	res := 0
	for _, acc := range accounts {
		sum := 0
		for _, n := range acc {
			sum += n
		}
		if res < sum {
			res = sum
		}
	}
	return res
}

func main() {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(accounts)
	fmt.Println(maximumWealth(accounts))
}
