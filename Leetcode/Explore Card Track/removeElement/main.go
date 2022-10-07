package main

import "fmt"

func removeElement(nums []int, val int) int {
	var counter int
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			if i < len(nums)-1 {
				copy(nums[i:], nums[i+1:])
			}
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]
		}
	}
	for j := range nums {
		if nums[j] != 0 {
			counter++
		}
	}
	return counter
}

func BetterremoveElement(nums []int, val int) int {
	count := 0
	for i, num := range nums {
		nums[i-count] = nums[i]
		if num == val {
			count += 1
		}
	}
	return len(nums) - count
}
func BestremoveElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
func MemoryBestremoveElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}
	return len(nums)
}
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
