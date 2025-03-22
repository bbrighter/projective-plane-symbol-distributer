package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"gonum.org/v1/gonum/stat/combin"
)

const NUMBER_OF_POINTS = 13

type ScoredBoard struct {
	Board                      Board
	Score                      int
	NumberOfColumnsWith3Colors int
	NumberOfColumnsWith4Colors int
}

func main() {
	symbols := NewSymbols()
	segments := []BaseSegment{
		{3, 6, 9, 0},
		{9, 7, 4, 1},
		{8, 9, 5, 2},
		{5, 4, 11, 3},
		{0, 10, 8, 4},
		{7, 0, 12, 5},
		{11, 8, 7, 6},
		{10, 3, 2, 7},
		{1, 12, 3, 8},
		{12, 11, 10, 9},
		{6, 5, 1, 10},
		{2, 1, 0, 11},
		{4, 2, 6, 12},
	}

	var start, end time.Time
	start = time.Now()

	topBoards := make([]ScoredBoard, 0, 5)

	redPossibilities := NUMBER_OF_POINTS - 1
	greenPossibilities := redPossibilities - 3
	yellowPossibilities := greenPossibilities - 3
	println("Total combinations: ", combin.Binomial(redPossibilities, 3)*combin.Binomial(greenPossibilities, 3)*combin.Binomial(yellowPossibilities, 3))
	redCombinations := combin.NewCombinationGenerator(redPossibilities, 3) // Ignore the joker! It will always be place on idx 12
	var dst []int
	i := 0
	redSymbols := symbols.Color(Red)
	greenSymbols := symbols.Color(Green)
	yellowSymbols := symbols.Color(Yellow)
	for redCombinations.Next() {
		rc := redCombinations.Combination(dst)
		redSymbolMap := redSymbols.ColorMap(rc)

		board := NewBoard(segments)
		board.SetColors(redSymbolMap)
		if !board.Validate(Red) {
			continue
		}
		if board.ColorCount(Red) != 12 {
			log.Panicf("Color count of red is %d, not 12", board.ColorCount(Red))
		}
		redMapping := createBijection(rc, NUMBER_OF_POINTS)
		greenCombinations := combin.NewCombinationGenerator(greenPossibilities, 3)
		for greenCombinations.Next() {
			gc := greenCombinations.Combination(dst)
			greenSymbolMap := greenSymbols.ColorMap(gc)
			board := NewBoard(segments)
			board.SetColors(redSymbolMap)
			board.RemapIndizes(redMapping)
			board.SetColors(greenSymbolMap)
			if board.ColorCount(Green) != 12 || board.ColorCount(Red) != 12 {
				log.Panicf("Color count of green is %d, of red is %d, not 12", board.ColorCount(Green), board.ColorCount(Red))
			}
			if !board.Validate(Green) {
				continue
			}
			greenMapping := createBijection(gc, NUMBER_OF_POINTS)
			yellowCombinations := combin.NewCombinationGenerator(yellowPossibilities, 3)
			for yellowCombinations.Next() {
				yc := yellowCombinations.Combination(dst)
				yellowSymbolMap := yellowSymbols.ColorMap(yc)

				board := NewBoard(segments)
				board.SetColors(redSymbolMap)
				board.RemapIndizes(redMapping)
				board.SetColors(greenSymbolMap)
				board.RemapIndizes(greenMapping)
				board.SetColors(yellowSymbolMap)
				if board.ColorCount(Green) != 12 || board.ColorCount(Red) != 12 || board.ColorCount(Yellow) != 12 {
					log.Panicf("Color count of green is %d, of red is %d, of yellow %d, not 12", board.ColorCount(Green), board.ColorCount(Red), board.ColorCount(Yellow))
				}
				if !board.Validate(Yellow) {
					continue
				}
				board.ColorUncoloredSegments(Blue)
				// Set joker at 13
				board.SetColor(12, Purple)
				if !board.Validate(Blue) {
					continue
				}
				i++

				scoredBoard := board.Evaluate()
				if scoredBoard.NumberOfColumnsWith4Colors == 4 {
					continue
				}
				if len(topBoards) < 5 {
					topBoards = append(topBoards, scoredBoard)
				} else {
					sort.Slice(topBoards, func(i, j int) bool {
						return topBoards[i].Score > topBoards[j].Score
					})
					if scoredBoard.Score > topBoards[len(topBoards)-1].Score {
						topBoards[len(topBoards)-1] = scoredBoard
					}
				}
			}
		}

	}
	println("Valid combinations: ", i)
	end = time.Now()

	for _, board := range topBoards {
		fmt.Println("---------------------------------------------------------------")
		board.Board.Print()
		fmt.Printf("--Score: %d-----3 colors: %d----------4 colors: %d----------- \n",
			board.Score,
			board.NumberOfColumnsWith3Colors,
			board.NumberOfColumnsWith4Colors, // FIX: why is this 0????
		)
	}

	fmt.Printf("Time taken: %s  \n", end.Sub(start))
}
