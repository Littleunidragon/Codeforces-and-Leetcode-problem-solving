package main

import "fmt"

func numberOfSteps(num int) int {
	var counter int
	for num != 0 {
		if num|1 > num {
			num = num / 2
			counter++
		} else {
			num = num - 1 //n--
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(numberOfSteps(123))
}
