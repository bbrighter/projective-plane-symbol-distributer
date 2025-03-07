package main

import (
	"fmt"
	"time"

	"modernc.org/mathutil"
)

func main() {
	symbols := NewSymbols()

	numbers := Numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	var start, end time.Time
	start = time.Now()

	i := 0
	// var wg sync.WaitGroup
	for mathutil.PermutationNext(numbers) {
		// wg.Add(1)
		i++
		if i > 10 {
			break
		}
		// go func(numbers Numbers) {
		symbolMap := symbols.ToMap(numbers)
		board := NewBoard(symbolMap)
		fmt.Println("---------------------------------------------------------------")
		board.Print()
		fmt.Println("---------------------------------------------------------------")
		// }(numbers)
	}
	// wg.Wait()
	end = time.Now()

	fmt.Printf("Time taken: %s  \n", end.Sub(start))
}
