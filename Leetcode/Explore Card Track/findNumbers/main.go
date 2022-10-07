package main

import "fmt"

func findNumbers(nums []int) int {
	var count int
	for i := range nums {
		s := fmt.Sprint(nums[i])
		if len(s)%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	s := []int{23, 456, 7}
	fmt.Println(findNumbers(s))
}
