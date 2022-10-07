package main

import "fmt"

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}

// function to get the absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// main function to get the sorted squares array
func BettersortedSquares(nums []int) []int {

	n := len(nums)           // get the length of the given array
	result := make([]int, n) // create a new slice of type int that can hold n values

	// left and right pointers
	left, right := 0, n-1

	// iterate through our new slice (result) in reverse as we will be adding numbers from largest to smallest
	for i := n - 1; i >= 0; i-- {
		square := 0 // the variable that will store the largest absolute value of the two pointers

		// if the absolute value of the left pointer is less than the absolute value of the right pointer,
		// set square to the larger number from the right pointer and move the right pointer one left.
		// Else set square to the larger number from the left pointer and move the left pointer one right.
		if Abs(nums[left]) < Abs(nums[right]) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}

		// square the value and add it to the appropriate index
		result[i] = square * square
	}

	// return the sorted array with the squared values
	return result
}

func main() {
	s := []int{-5, -3, 0, 7}
	fmt.Println(sortedSquares(s))
}
