package main

import "fmt"

func maxSubarrayLength(nums []int, k int) int {
	size := len(nums)
	start, end := 0, 0
	freq := make(map[int]int)
	maxLen := 0
	for start <= end && end < size {
		if freq[nums[end]] < k {
			freq[nums[end]]++
			end++
			if end-start > maxLen {
				maxLen = end - start
			}
			continue
		}
		for i := start; i <= end; i++ {
			if nums[i] == nums[end] {
				start = i + 1
				freq[nums[end]]--
				break
			} else {
				freq[nums[i]]--
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Println(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))
}
