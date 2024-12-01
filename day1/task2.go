package main

import (
	"fmt"
)

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	var left []int
	var right []int
	for {
		var a int
		var b int

		num, _ := fmt.Scanln(&a, &b)
		if num != 2 {
			break
		}

		left = append(left, a)
		right = append(right, b)
	}

	simScore := 0
	for i := 0; i < len(left); i++ {
		rightCount := 0
		for _, el := range right {
			if el == left[i] {
				rightCount++
			}
		}

		simScore += left[i] * rightCount
	}

	println(simScore)
}
