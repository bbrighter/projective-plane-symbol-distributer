package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	symbols := NewSymbols()
	symbolMap := symbols.ToMap(Numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})

	board := NewBoard(symbolMap)
	assert.Len(t, board.Segments, 13)
	assert.Equal(t, Symbol{Color: Red, Shape: Circle}, board.Segments[0].Inner)
}
