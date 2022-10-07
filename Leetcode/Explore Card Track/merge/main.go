package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	k := m + n - 1
	i := m - 1
	j := n - 1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}
func Bettermerge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

func main() {
	nums1 := []int{0, 0, 0}
	n := 0
	nums2 := []int{2, 5, 6}
	m := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
