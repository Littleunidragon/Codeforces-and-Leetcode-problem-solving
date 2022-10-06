package main

func runningSum(nums []int) []int {
	runningsum := make([]int, len(nums))
	runningsum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		runningsum[i] = runningsum[i-1] + nums[i]
	}
	return runningsum
}
