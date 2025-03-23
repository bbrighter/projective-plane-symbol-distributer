package main

import "fmt"

type Board struct {
	Segments []Segment
}

func NewBoard(baseSegments []BaseSegment) Board {
	var segments []Segment
	for _, bs := range baseSegments {
		var segment Segment
		segment.Inner.Index = bs.Inner
		segment.Middle.Index = bs.Middle
		segment.Outer.Index = bs.Outer
		segment.Outside.Index = bs.Outside
		segments = append(segments, segment)
	}

	return Board{Segments: segments}
}

func (b *Board) SetColors(colors map[int]Color) {
	for index, symbol := range colors {
		b.SetColor(index, symbol)
	}
}

func (b *Board) SetColor(index int, color Color) {
	for i, s := range b.Segments {
		if s.Inner.Index == index {
			b.Segments[i].Inner.Color = color
		}
		if s.Middle.Index == index {
			b.Segments[i].Middle.Color = color
		}
		if s.Outer.Index == index {
			b.Segments[i].Outer.Color = color
		}
		if s.Outside.Index == index {
			b.Segments[i].Outside.Color = color
		}
	}
}

func (b *Board) SetShapes(indexes []int) {
	for i, idx := range indexes {
		var useShape Shape
		switch i {
		case 0:
			useShape = Circle
		case 1:
			useShape = Square
		case 2:
			useShape = Triangle
		}
		b.SetShape(idx, useShape)
	}
}

func (b *Board) SetShape(index int, shape Shape) {
	for i, s := range b.Segments {
		if s.Inner.Index == index {
			b.Segments[i].Inner.Shape = shape
		}
		if s.Middle.Index == index {
			b.Segments[i].Middle.Shape = shape
		}
		if s.Outer.Index == index {
			b.Segments[i].Outer.Shape = shape
		}
		if s.Outside.Index == index {
			b.Segments[i].Outside.Shape = shape
		}
	}
}

func (b Board) GetJokerIndex() int {
	jokerIndex := 0
	for i, s := range b.Segments {
		if s.Inner.Color == 0 {
			jokerIndex = b.Segments[i].Inner.Index
			break
		}
	}
	return jokerIndex
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
	numberOfColumnsWith3Colors := 0
	numberOfColumnsWith4Colors := 0
	for _, s := range b.Segments {
		boardScore, uniqueColors := s.Evaluate()
		score += boardScore
		if uniqueColors == 3 {
			numberOfColumnsWith3Colors++
		}
		if uniqueColors == 4 {
			numberOfColumnsWith4Colors++
		}
	}
	return ScoredBoard{
		Board:                      b,
		Score:                      score,
		NumberOfColumnsWith3Colors: numberOfColumnsWith3Colors,
		NumberOfColumnsWith4Colors: numberOfColumnsWith4Colors,
	}
}

func (b Board) Validate(color Color) bool {
	for _, s := range b.Segments {
		if !s.Validate(color) {
			return false
		}
	}
	return true
}

func (b Board) GetColorIndexes(color Color) []int {
	indexes := make(map[int]bool)
	for _, s := range b.Segments {
		if s.Inner.Color == color {
			indexes[s.Inner.Index] = true
		}
		if s.Middle.Color == color {
			indexes[s.Middle.Index] = true
		}
		if s.Outer.Color == color {
			indexes[s.Outer.Index] = true
		}
		if s.Outside.Color == color {
			indexes[s.Outside.Index] = true
		}
	}
	var result []int
	for key := range indexes {
		result = append(result, key)
	}
	return result
}

func (b Board) ColorCount(color Color) int {
	count := 0
	for _, s := range b.Segments {
		if s.Inner.Color == color {
			count++
		}
		if s.Middle.Color == color {
			count++
		}
		if s.Outer.Color == color {
			count++
		}
		if s.Outside.Color == color {
			count++
		}
	}
	return count
}

func (b *Board) RemapIndizes(fromTo map[int]int) {
	for i, s := range b.Segments {
		b.Segments[i].Inner.Index = fromTo[s.Inner.Index]
		b.Segments[i].Middle.Index = fromTo[s.Middle.Index]
		b.Segments[i].Outer.Index = fromTo[s.Outer.Index]
		b.Segments[i].Outside.Index = fromTo[s.Outside.Index]
	}
}

func (b *Board) ColorUncoloredSegments(color Color) {
	for i, s := range b.Segments {
		if s.Inner.Color == 0 {
			b.Segments[i].Inner.Color = color
		}
		if s.Outer.Color == 0 {
			b.Segments[i].Outer.Color = color
		}
		if s.Middle.Color == 0 {
			b.Segments[i].Middle.Color = color
		}
		if s.Outside.Color == 0 {
			b.Segments[i].Outside.Color = color
		}
	}
}
