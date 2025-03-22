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

func (s Segment) Evaluate() (int, int) {
	colorSet := make(map[Color]struct{})
	shapeSet := make(map[Shape]struct{})

	symbols := []Symbol{s.Inner, s.Middle, s.Outer, s.Outside}
	for _, symbol := range symbols {
		colorSet[symbol.Color] = struct{}{}
		shapeSet[symbol.Shape] = struct{}{}
	}
	uniqueColors := len(colorSet)
	var uniqueColorPoints int
	switch uniqueColors {
	case 4:
		uniqueColorPoints = 100
	case 3:
		uniqueColorPoints = 0
	default:
		uniqueColorPoints = -50
	}

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

	return uniqueColorPoints + uniqueShapes*0 - colorPenalty*1 - shapePenalty*0, uniqueColors
}

// A segment is not valid if there are only 2 different colors
func (s Segment) Validate(color Color) bool {
	colorCount := 0

	symbols := []Symbol{s.Inner, s.Middle, s.Outer, s.Outside}
	for _, symbol := range symbols {
		if symbol.HasColor(color) {
			colorCount++
		}
	}
	return colorCount < 3
}
