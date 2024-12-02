package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const maxDist = 3

func removeAt(sl []int, index int) []int {
	var newSlice []int
	for ind, elem := range sl {
		if ind == index {
			continue
		}
		newSlice = append(newSlice, elem)
	}
	return newSlice
}

func isReportSafe(levels []int, skipped bool) bool {
	increasing := false
	prev := 0
	current := 0
	for ind, elem := range levels {
		prev = current
		current = elem

		if ind == 0 {
			continue
		}

		if ind == 1 {
			if prev < current && prev+maxDist >= current {
				increasing = true
			} else if prev > current && prev-maxDist <= current {
				increasing = false
			} else {
				if !skipped {
					if isReportSafe(removeAt(levels, ind-1), true) {
						return true
					}
					if isReportSafe(removeAt(levels, ind), true) {
						return true
					}
				}
				return false
			}
			continue
		}

		if !(increasing && prev < current && prev+maxDist >= current) &&
			!(!increasing && prev > current && prev-maxDist <= current) {
			if !skipped {
				if isReportSafe(removeAt(levels, ind-2), true) {
					return true
				}
				if isReportSafe(removeAt(levels, ind-1), true) {
					return true
				}
				if isReportSafe(removeAt(levels, ind), true) {
					return true
				}
			}
			return false
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	safeCount := 0

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\r\n")
		if err != nil {
			break
		}

		var levels []int
		elems := strings.Split(line, " ")

		for _, elem := range elems {
			level, _ := strconv.Atoi(elem)
			levels = append(levels, level)
		}

		if isReportSafe(levels, false) {
			safeCount++
		}
	}

	println(safeCount)
}
