package main

import "fmt"

type Segment struct {
	Inner   Symbol
	Middle  Symbol
	Outer   Symbol
	Outside Symbol
}

type Board struct {
	Segments [13]Segment
}

func NewBoard(symbols map[int]Symbol) Board {
	base := NewBaseBoard()

	var segments [13]Segment
	for i, segment := range base.Segments {
		segments[i].Inner = symbols[segment.Inner]
		segments[i].Middle = symbols[segment.Middle]
		segments[i].Outer = symbols[segment.Outer]
		segments[i].Outside = symbols[segment.Outside]
	}
	return Board{Segments: segments}
}

func (b Board) Print() {
	var topString, middleString, bottomString, outerString string
	for _, segment := range b.Segments {
		outerString += " " + segment.Outside.String() + " |"
		topString += " " + segment.Outer.String() + " |"
		middleString += " " + segment.Middle.String() + " |"
		bottomString += " " + segment.Inner.String() + " |"
	}
	fmt.Println("| " + outerString)
	fmt.Println("| " + topString)
	fmt.Println("| " + middleString)
	fmt.Println("| " + bottomString)
}
