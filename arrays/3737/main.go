package main

import "fmt"

func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	mapped := make([]int, n)
	for i, v := range nums {
		if v == target {
			mapped[i] = 1
		} else {
			mapped[i] = -1
		}
	}
	result := 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += mapped[j]
			if sum > 0 {
				result++
			}

		}
	}

	return result
}

func main() {
	fmt.Println(countMajoritySubarrays([]int{1, 2, 2, 3}, 2))
}
