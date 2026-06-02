package main

import (
	"fmt"
	"math"
)

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	ans1 := math.MaxInt
	ans2 := math.MaxInt

	// water + land
	wl := len(waterStartTime)
	ll := len(landStartTime)
	wcl := waterStartTime[0] + waterDuration[0]

	for i := 0; i < wl; i++ {
		wcl = min(wcl, waterStartTime[i]+waterDuration[i])
	}

	for j := 0; j < ll; j++ {
		if landStartTime[j] <= wcl {
			ans1 = min(ans1, wcl+landDuration[j])
		} else {
			ans1 = min(ans1, landDuration[j]+landStartTime[j])
		}

	}

	// land + water
	lcl := landStartTime[0] + landDuration[0]
	for i := 0; i < ll; i++ {
		lcl = min(lcl, landStartTime[i]+landDuration[i])
	}
	for j := 0; j < wl; j++ {
		if waterStartTime[j] <= lcl {
			ans2 = min(ans2, lcl+waterDuration[j])
		} else {
			ans2 = min(ans2, waterDuration[j]+waterStartTime[j])
		}
	}

	return min(ans1, ans2)

}

func main() {
	fmt.Println(earliestFinishTime([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
}
