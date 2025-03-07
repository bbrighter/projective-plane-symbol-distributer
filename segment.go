package main

type BaseSegment struct {
	Inner   int
	Middle  int
	Outer   int
	Outside int
}

type Segment struct {
	Inner   Symbol
	Middle  Symbol
	Outer   Symbol
	Outside Symbol
}

func (s Segment) Evaluate() int {
	colorSet := make(map[Color]struct{})
	shapeSet := make(map[Shape]struct{})

	symbols := []Symbol{s.Inner, s.Middle, s.Outer, s.Outside}
	for _, symbol := range symbols {
		colorSet[symbol.Color] = struct{}{}
		shapeSet[symbol.Shape] = struct{}{}
	}
	uniqueColors := len(colorSet)
	uniqueShapes := len(shapeSet)

	var colorPenalty, shapePenalty int = 0, 0
	for i := 1; i < len(symbols); i++ {
		if symbols[i].Color == symbols[i-1].Color {
			colorPenalty++
		}
		if symbols[i].Shape == symbols[i-1].Shape {
			shapePenalty++
		}
	}

	return uniqueColors*5 + uniqueShapes*3 - colorPenalty*4 - shapePenalty*2
}
