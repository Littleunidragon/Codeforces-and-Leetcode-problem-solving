package main

func checkIfExist(arr []int) bool {
	for j := 0; j != len(arr); j++ {
		if arr[j]%2 == 0 {
			for i := 0; i < len(arr); i++ {
				if i == j {
					continue
				}
				if arr[j]*2 == arr[i] {
					return true
				}
			}
		}

	}
	return false
}

func EasiercheckIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == 2*arr[j] {
				return true
			}
		}
	}
	return false
}
func MapcheckIfExist(arr []int) bool {
	m := map[int]int{}
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			continue
		}
		if c, ok := m[arr[i]/2]; ok {
			if arr[i] != 0 {
				return true
			}
			if c > 1 {
				return true
			}
		}
	}
	return false
}
func main() {

}
