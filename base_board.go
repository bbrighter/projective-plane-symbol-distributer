package main

import "fmt"

type BaseSegment struct {
	Inner   int
	Middle  int
	Outer   int
	Outside int
}

type BaseBoard struct {
	Segments []BaseSegment
}

func NewBaseBoard() BaseBoard {
	return BaseBoard{
		Segments: []BaseSegment{
			{4, 7, 10, 1},
			{10, 8, 5, 2},
			{9, 10, 6, 3},
			{6, 5, 12, 4},
			{1, 11, 9, 5},
			{8, 1, 13, 6},
			{12, 9, 8, 7},
			{11, 4, 3, 8},
			{2, 13, 4, 9},
			{13, 12, 11, 10},
			{7, 6, 2, 11},
			{3, 2, 1, 12},
			{5, 3, 7, 13},
		},
	}
}

func (b *BaseBoard) Print() {
	var topString, middleString, bottomString, outerString string
	for _, segment := range b.Segments {
		outerString += " " + fmt.Sprintf("(%2d)", segment.Outside) + " |"
		topString += " " + fmt.Sprintf("%4d", segment.Outer) + " |"
		middleString += " " + fmt.Sprintf("%4d", segment.Middle) + " |"
		bottomString += " " + fmt.Sprintf("%4d", segment.Inner) + " |"
	}
	fmt.Println(outerString)
	fmt.Println(topString)
	fmt.Println(middleString)
	fmt.Println(bottomString)
}
