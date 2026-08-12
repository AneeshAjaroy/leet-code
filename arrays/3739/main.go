package main

import (
	"fmt"
	"slices"
)

func countMajoritySubarrays(nums []int, target int) int64 {

	// create prefix sum array
	n := len(nums)
	prefixSum := make([]int, n+1)
	prefixSum[0] = 0
	for i, v := range nums {
		if v == target {
			prefixSum[i+1] = prefixSum[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i] - 1
		}

	}

	// map the prefixes to natural numbers
	// first sort them
	dummyPrefixSum := slices.Clone(prefixSum)
	slices.Sort(dummyPrefixSum)
	mapperVal := 1
	uniqueMap := make(map[int]int)
	for _, v := range dummyPrefixSum {
		if uniqueMap[v] == 0 {
			uniqueMap[v] = mapperVal
			mapperVal++
		}
	}

	// now build and use fenwick tree
	// it is just an array with some special properties
	var result int64
	fenwick := make([]int, mapperVal)

	for _, v := range prefixSum {
		result += int64(query(fenwick, uniqueMap[v]-1))
		update(fenwick, uniqueMap[v])
	}
	return result
}

func update(arr []int, i int) {
	for i < len(arr) {
		arr[i] += 1
		i += i & -i
	}
}

func query(arr []int, i int) int {
	sum := 0
	for i > 0 {
		sum += arr[i]
		i -= i & -i
	}
	return sum
}

func main() {
	fmt.Println(countMajoritySubarrays([]int{1, 2, 2, 3}, 2))
}
