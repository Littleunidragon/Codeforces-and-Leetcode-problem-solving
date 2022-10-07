package main

func findMaxConsecutiveOnes(nums []int) int {
	var count, ans int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}

func main() {

}
