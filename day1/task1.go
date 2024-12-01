package main

import (
	"fmt"
	"sort"
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

	sort.Ints(left)
	sort.Ints(right)

	diffScore := 0
	for i := 0; i < len(left); i++ {
		diffScore += Abs(left[i] - right[i])
	}

	println(diffScore)
}
