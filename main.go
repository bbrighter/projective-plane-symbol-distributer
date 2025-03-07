package main

import (
	"fmt"
	"sort"
	"time"

	"modernc.org/mathutil"
)

type ScoredBoard struct {
	Board Board
	Score int
}

func main() {
	symbols := NewSymbols()
	segments := []BaseSegment{{4, 7, 10, 1},
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
		{5, 3, 7, 13}}

	numbers := Numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	var start, end time.Time
	start = time.Now()

	i := 0
	topBoards := make([]ScoredBoard, 0, 5)
	// var wg sync.WaitGroup
	for mathutil.PermutationNext(numbers) {
		// wg.Add(1)
		i++
		startFrom := 0
		if i < startFrom {
			continue
		}
		if i > startFrom+1_000_000 {
			break
		}
		// go func(numbers Numbers) {
		symbolMap := symbols.ToMap(numbers)
		board := NewBoard(symbolMap, segments)
		scoredBoard := board.Evaluate()
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

		// }(numbers)
	}
	// wg.Wait()
	end = time.Now()

	for _, board := range topBoards {
		fmt.Println("---------------------------------------------------------------")
		board.Board.Print()
		fmt.Printf("------------------Score: %d---------------------------------------- \n", board.Score)
	}

	fmt.Printf("Time taken: %s  \n", end.Sub(start))
}
