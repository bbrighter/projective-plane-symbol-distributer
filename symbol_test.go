package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSymbols(t *testing.T) {
	t.Parallel()
	symbols := NewSymbols()

	assert.Len(t, symbols, 13)
	assert.Contains(t, symbols, Symbol{Color: Red, Shape: Square})
	assert.Contains(t, symbols, Symbol{Color: Purple, Shape: Star})
}

func TestToMap(t *testing.T) {
	t.Parallel()

	symbols := NewSymbols()
	assert.Equal(t, Symbol{Color: Red, Shape: Circle}, symbols[0])

	permutation := Numbers{2, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	mapping := symbols.ToMap(permutation)

	assert.Len(t, mapping, 13)
	assert.Equal(t, Symbol{Color: Red, Shape: Circle}, mapping[2], "First symbol is now at second number")
}
