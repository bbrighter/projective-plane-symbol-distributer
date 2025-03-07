package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	t.Parallel()

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

			board := NewBoard(symbolMap)
			assert.Len(t, board.Segments, 13)
			assert.Equal(t, test.expectedFirstInner, board.Segments[0].Inner)
		})
	}
}
