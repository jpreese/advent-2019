package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const min = 245318
	const max = 765747

	var firstTotal int
	for i := min; i < max; i++ {
		input := strconv.Itoa(i)
		if runFirstRules(input) {
			firstTotal++
		}
	}
	fmt.Println(firstTotal)

	var secondTotal int
	for i := min; i < max; i++ {
		input := strconv.Itoa(i)
		if runSecondRules(input) {
			secondTotal++
		}
	}
	fmt.Println(secondTotal)
}

func runFirstRules(input string) bool {
	var hasDouble bool
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			return false
		}

		if input[i] == input[i+1] {
			hasDouble = true
		}
	}

	return hasDouble
}

func runSecondRules(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			return false
		}
	}

	slice := strings.Split(input, "")
	sort.Strings(slice)

	var hasDouble bool
	for i := 0; i < 10; i++ {
		var numberCount int
		for _, digit := range slice {
			if digit == strconv.Itoa(i) {
				numberCount++
			}
		}

		if numberCount == 2 {
			hasDouble = true
		}
	}

	return hasDouble
}
