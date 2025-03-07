package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateSegment(t *testing.T) {

	tests := map[string]struct {
		segment       Segment
		expectedScore int
	}{
		"highest score": {
			Segment{
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Blue, Shape: Square},
				Symbol{Color: Green, Shape: Triangle},
				Symbol{Color: Purple, Shape: Star},
			}, 256},
		"lowest score": {
			Segment{
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Red, Shape: Circle},
			}, 1},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			score := test.segment.Evaluate()

			assert.Equal(t, test.expectedScore, score)
		})
	}
}
