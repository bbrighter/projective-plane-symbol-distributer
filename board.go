package main

import "fmt"

type Segment struct {
	Inner   Symbol
	Middle  Symbol
	Outer   Symbol
	Outside Symbol
}

type Board struct {
	Segments []Segment
}

func NewBoard(symbols map[int]Symbol) Board {
	base := NewBaseBoard()

	var segments []Segment
	for _, bs := range base.Segments {
		var segment Segment
		segment.Inner = symbols[bs.Inner]
		segment.Middle = symbols[bs.Middle]
		segment.Outer = symbols[bs.Outer]
		segment.Outside = symbols[bs.Outside]
		segments = append(segments, segment)
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
