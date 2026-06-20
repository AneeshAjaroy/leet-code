package main

import (
	"fmt"
	"slices"
)

// OOM
// func maxBuilding(n int, restrictions [][]int) int {
// 	height := make([]int, n)
// 	set := make([]bool, n)
// 	height[0] = 0
// 	set[0] = true
// 	height[1] = 1
// 	set[1] = true
// 	maxHeight := 1
// 	for _, res := range restrictions {
// 		height[res[0]-1] = res[1]
// 		set[res[0]-1] = true
// 	}
// 	for i := 1; i < n; {
// 		if set[i] {
// 			if height[i] > height[i-1] {
// 				if height[i]-height[i-1] <= 1 {
// 					i++
// 				} else {
// 					height[i] = height[i-1] + 1
// 					i++
// 				}
// 			} else {
// 				if height[i-1]-height[i] <= 1 {
// 					i++
// 				} else {
// 					height[i-1] = height[i] + 1
// 					i--
// 				}
// 			}
// 		} else {
// 			height[i] = height[i-1] + 1
// 			set[i] = true
// 		}
// 	}

// 	maxHeight = 0
// 	for _, v := range height {
// 		if v > maxHeight {
// 			maxHeight = v
// 		}
// 	}
// 	return maxHeight
// }

func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})
	slices.SortFunc[[][]int](restrictions, func(i, j []int) int {
		return i[0] - j[0]
	})
	numRest := len(restrictions)
	for i := 1; i < numRest; i++ {
		restrictions[i][1] = min(restrictions[i][1], restrictions[i][0]-restrictions[i-1][0]+restrictions[i-1][1])
	}
	for i := numRest - 2; i > 0; i-- {
		restrictions[i][1] = min(restrictions[i][1], restrictions[i+1][0]-restrictions[i][0]+restrictions[i+1][1])
	}

	maxHeight := 0
	for i := 1; i < numRest; i++ {
		maxHeight = max(maxHeight, (restrictions[i][1]+restrictions[i-1][1]+restrictions[i][0]-restrictions[i-1][0])/2)
	}
	return max(maxHeight, restrictions[numRest-1][1]+n-restrictions[numRest-1][0])

}

func main() {
	fmt.Println(maxBuilding(5, [][]int{{2, 1}, {4, 1}}))
}
