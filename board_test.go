package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	t.Parallel()

	segments := []BaseSegment{{4, 7, 10, 1},
		{10, 8, 5, 2},
		{9, 10, 6, 3},
		{6, 5, 12, 4},
		{1, 11, 9, 5},
		{8, 1, 13, 6},
		{12, 9, 8, 7},
		{11, 4, 3, 8},
		{2, 13, 4, 9},
		{13, 12, 11, 10},
		{7, 6, 2, 11},
		{3, 2, 1, 12},
		{5, 3, 7, 13}}

	tests := map[string]struct {
		permutation        Numbers
		expectedFirstInner Symbol
	}{
		"no permutation": {
			permutation:        Numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			expectedFirstInner: Symbol{Color: Green, Shape: Circle},
		},
		"(14) permutation": {
			permutation:        Numbers{4, 2, 3, 1, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			expectedFirstInner: Symbol{Color: Red, Shape: Circle},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			symbols := NewSymbols()
			symbolMap := symbols.ToMap(test.permutation)

			board := NewBoard(symbolMap, segments)
			assert.Len(t, board.Segments, 13)
			assert.Equal(t, test.expectedFirstInner, board.Segments[0].Inner)
		})
	}
}
