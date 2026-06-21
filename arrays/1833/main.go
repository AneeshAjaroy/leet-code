package main

import "fmt"

func maxIceCream(costs []int, coins int) int {
	// find max
	maxVal := 0
	for _, v := range costs {
		if v > maxVal {
			maxVal = v
		}
	}
	// fill freq arr
	freqArr := make([]int, maxVal+1)
	for _, v := range costs {
		freqArr[v]++
	}

	// max Ice Creams can be bought
	maxBought := 0
	cost := 0
	for i, v := range freqArr {
		if cost+i*v <= coins {
			maxBought += v
			cost += i * v
		} else {
			maxBought += (coins - cost) / i
			break

		}
	}
	return maxBought

}

func main() {
	fmt.Println(maxIceCream([]int{2, 2, 2, 2}, 5))
}
