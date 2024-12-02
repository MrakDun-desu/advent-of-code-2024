package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const maxDist = 3

func main() {
	reader := bufio.NewReader(os.Stdin)
	safeCount := 0

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			break
		}
		elems := strings.Split(line, " ")
		increasing := false
		prev := 0
		current := 0

		for ind, elem := range elems {
			prev = current
			current, _ = strconv.Atoi(elem)

			if ind == 0 {
				continue
			}

			if ind == 1 {
				if prev < current && prev+maxDist >= current {
					increasing = true
				} else if prev > current && prev-maxDist <= current {
					increasing = false
				} else {
					break
				}
				continue
			}

			if (increasing && prev < current && prev+maxDist >= current) ||
				(!increasing && prev > current && prev-maxDist <= current) {
				if ind+1 == len(elems) {
					safeCount++
				}
			} else {
				break
			}
		}
	}

	println(safeCount)
}
