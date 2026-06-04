package main

import (
	"fmt"
	"slices"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	slices.Sort(asteroids)
	for _, a := range asteroids {
		if a > mass {
			return false
		} else {
			mass += a
		}
	}
	return true
}

func main() {
	fmt.Println(asteroidsDestroyed(5, []int{4, 9, 23, 4}))
}
