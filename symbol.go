package main

import "fmt"

type Color int

const (
	Red Color = iota
	Green
	Blue
	Yellow
	Purple
)

type Shape string

const (
	Circle   Shape = "○"
	Square   Shape = "□"
	Triangle Shape = "△"
	Star     Shape = "X"
)

type Symbol struct {
	Color Color
	Shape Shape
}

type Symbols [13]Symbol

var Colors = []Color{Red, Green, Blue, Yellow}
var Shapes = []Shape{Circle, Square, Triangle}

func NewSymbols() Symbols {
	var symbols [13]Symbol
	i := 0
	for _, color := range Colors {
		for _, shape := range Shapes {
			symbols[i] = Symbol{Color: color, Shape: shape}
			i++
		}
	}
	symbols[12] = Symbol{Color: Purple, Shape: Star}
	return symbols
}

func (symbols Symbols) ToMap(permutation Numbers) map[int]Symbol {
	var symbolMap = make(map[int]Symbol)
	for index, value := range permutation {
		symbolMap[value] = symbols[index]
	}
	return symbolMap
}

func (s Symbol) String() string {

	var colorCodes = map[Color]string{
		Red:    "\033[31m",
		Green:  "\033[32m",
		Blue:   "\033[34m",
		Yellow: "\033[33m",
		Purple: "\033[35m",
	}

	const reset = "\033[0m"

	colorCode, exists := colorCodes[s.Color]
	if !exists {
		colorCode = reset
	}
	return fmt.Sprintf("%s%s%s", colorCode, s.Shape, reset)

}
