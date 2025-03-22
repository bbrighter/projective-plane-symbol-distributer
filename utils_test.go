package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapForLessNumbers(t *testing.T) {
	tests := map[string]struct {
		totalNumbers   int
		numbersInUse   []int
		expectedResult map[int]int
	}{
		"0 in use": {
			totalNumbers:   2,
			numbersInUse:   []int{0},
			expectedResult: map[int]int{0: 1, 1: 0},
		},
		"0 and 1 in use, total 4": {
			totalNumbers:   4,
			numbersInUse:   []int{0, 1},
			expectedResult: map[int]int{0: 2, 1: 3, 2: 0, 3: 1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := createBijection(test.numbersInUse, test.totalNumbers)
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
