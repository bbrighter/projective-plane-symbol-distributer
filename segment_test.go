package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateSegment(t *testing.T) {

	tests := map[string]struct {
		segment                      Segment
		expectedScore                int
		expectedNumberOfUniqueColors int
	}{
		"highest score": {
			segment: Segment{
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Blue, Shape: Square},
				Symbol{Color: Green, Shape: Triangle},
				Symbol{Color: Purple, Shape: Star},
			},
			expectedScore:                100,
			expectedNumberOfUniqueColors: 4,
		},
		"lowest score": {
			segment: Segment{
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Red, Shape: Circle},
				Symbol{Color: Red, Shape: Circle},
			},
			expectedScore:                -53,
			expectedNumberOfUniqueColors: 1,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			score, uniqueColors := test.segment.Evaluate()

			assert.Equal(t, test.expectedNumberOfUniqueColors, uniqueColors)
			assert.Equal(t, test.expectedScore, score)
		})
	}
}
