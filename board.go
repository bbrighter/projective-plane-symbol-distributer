package main

import "fmt"

type Board struct {
	Segments []Segment
}

func NewBoard(symbols map[int]Symbol, baseSegments []BaseSegment) Board {
	var segments []Segment
	for _, bs := range baseSegments {
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

func (b Board) Evaluate() ScoredBoard {
	score := 0
	for _, s := range b.Segments {
		score += s.Evaluate()
	}
	return ScoredBoard{Board: b, Score: score}
}
