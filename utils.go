package main

import "slices"

func createBijection(numbersInUse []int, numberOfPoints int) map[int]int {
	var numbersNotInUse []int
	for val := range numberOfPoints {
		if !slices.Contains(numbersInUse, val) {
			numbersNotInUse = append(numbersNotInUse, val)
		}
	}

	result := make(map[int]int)
	for i, no := range numbersNotInUse {
		result[no] = i
	}
	for i, no := range numbersInUse {
		result[no] = i + len(numbersNotInUse)
	}
	return result
}

func applyPermutation(input []int, permutation []int) []int {
	result := make([]int, len(input))
	for i, p := range permutation {
		result[i] = input[p]
	}
	return result
}
