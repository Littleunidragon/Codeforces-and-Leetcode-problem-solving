package main

func removeDuplicates(nums []int) int {
	for i := 0; i != len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}
	return len(nums)
}

func main() {

}
