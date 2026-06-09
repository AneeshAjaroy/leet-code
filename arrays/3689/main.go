package main

func maxTotalValue(nums []int, k int) int64 {
	maxVal := nums[0]
	minVal := nums[0]
	for _, val := range nums {
		if maxVal < val {
			maxVal = val
		}
		if minVal > val {
			minVal = val
		}
	}
	val := maxVal - minVal
	val64 := int64(val)
	k64 := int64(k)

	return val64 * k64
}
