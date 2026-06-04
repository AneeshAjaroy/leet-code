package main

import (
	"fmt"
	"strconv"
)

func totalWaviness(num1 int, num2 int) int {
	ans := 0
	for i := num1; i <= num2; i++ {
		str := strconv.Itoa(i)
		n := len(str)
		if n < 3 {
			continue
		}
		for j := 1; j < n-1; j++ {
			if (str[j] < str[j-1] && str[j] < str[j+1]) || (str[j] > str[j-1] && str[j] > str[j+1]) {
				ans += 1
			}
		}
	}
	return ans

}

func main() {
	fmt.Println(totalWaviness(4848, 4849))
}
