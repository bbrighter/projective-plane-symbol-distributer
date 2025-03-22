package main

import "fmt"

type Color int

const (
	Red Color = iota + 1
	Green
	Blue
	Yellow
	Purple
)

type Shape string

const (
	Circle   Shape = "○"
	Square   Shape = "□"
	Triangle Shape = "A"
	Star     Shape = "X"
)

type Symbol struct {
	Color Color
	Shape Shape
	Index int
}

type Symbols []Symbol

var Colors = []Color{Red, Green, Blue, Yellow}
var Shapes = []Shape{Circle, Square, Triangle}

func NewSymbols() Symbols {
	var symbols Symbols
	for _, color := range Colors {
		for _, shape := range Shapes {
			symbol := Symbol{Color: color, Shape: shape}
			symbols = append(symbols, symbol)
		}
	}
	symbols = append(symbols, Symbol{Color: Purple, Shape: Star})
	return symbols
}

func (symbols Symbols) ColorMap(permutation []int) map[int]Color {
	var symbolMap = make(map[int]Color)
	for index, value := range permutation {
		symbolMap[value] = symbols[index].Color
	}
	return symbolMap
}

func (symbols Symbols) Color(color Color) Symbols {
	var coloredSymbols Symbols
	for _, s := range symbols {
		if s.Color == color {
			coloredSymbols = append(coloredSymbols, s)
		}
	}
	return coloredSymbols
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
	if s.Shape == "" {
		s.Shape = "●"
	}
	return fmt.Sprintf("%s%s%s", colorCode, s.Shape, reset)

}

func (s Symbol) HasColor(color Color) bool {
	return s.Color == color
}

func (s Symbol) HasShape() bool {
	return s.Shape != ""
}
