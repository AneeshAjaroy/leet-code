package main

import (
	"fmt"
)

func angleClock(hour int, minutes int) float64 {
	var hourDeg, minDeg float64
	if hour == 12 {
		hourDeg = float64(minutes) * 0.5
	} else {
		hourDeg = float64(hour*60+minutes) * 0.5
	}
	minDeg = float64(minutes * 6)
	diff := minDeg - hourDeg
	if diff < 0 {
		diff = -diff
	}
	if diff > 180 {

		diff = 360 - diff
	}

	return diff
}

func main() {
	fmt.Println(angleClock(8, 7))
}
