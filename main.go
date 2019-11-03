package main

import "fmt"

const width int = 10
const height int = 10

func getCellState(actualState bool, neighboursCount int) bool {
	if actualState {
		return neighboursCount >= 2 && neighboursCount <= 3
	}
	return neighboursCount == 3
}

func countNeighbours(frame [height][width]bool, y int, x int) int {
	neighboursCount := 0
	if x-1 >= 0 && frame[y][x-1] {
		neighboursCount++
	}
	if x+1 < len(frame[y]) && frame[y][x+1] {
		neighboursCount++
	}
	if y-1 >= 0 {
		if x-1 >= 0 && frame[y-1][x-1] {
			neighboursCount++
		}
		if frame[y-1][x] {
			neighboursCount++
		}
		if x+1 < len(frame[y]) && frame[y-1][x+1] {
			neighboursCount++
		}
	}
	if y+1 < len(frame) {
		if x-1 >= 0 && frame[y+1][x-1] {
			neighboursCount++
		}
		if frame[y+1][x] {
			neighboursCount++
		}
		if x+1 < len(frame[y]) && frame[y+1][x+1] {
			neighboursCount++
		}
	}
	return neighboursCount
}

func getNextFrame(actualFrame [height][width]bool) [height][width]bool {
	var newFrame [height][width]bool
	for y := 0; y < len(actualFrame); y++ {
		for x := 0; x < len(actualFrame[0]); x++ {
			neighboursCount := countNeighbours(actualFrame, y, x)
			newFrame[y][x] = getCellState(actualFrame[y][x], neighboursCount)
		}
	}
	return newFrame
}

func printRow(row [width]bool) {
	for i := 0; i < len(row); i++ {
		if row[i] == true {
			fmt.Print("*")
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("\n")
}

func main() {
	frame := [height][width]bool{
		{false, true, false, false, true, false, true, false, true, true},
		{true, false, true, false, false, false, false, true, true, true},
		{false, false, true, false, true, false, false, true, false, true},
		{false, true, false, false, true, false, false, false, false, false},
		{false, true, false, false, true, false, false, false, false, false},
		{false, true, false, false, true, false, false, false, false, false},
		{true, false, true, false, false, false, false, true, true, true},
		{false, false, true, false, true, false, false, true, false, true},
		{false, true, false, false, true, false, false, false, false, false},
		{false, true, false, false, true, false, false, false, false, false},
	}
	for row := 0; row < height; row++ {
		printRow(frame[row])
	}
	for i := 0; i < 10000000000000*10000; i++ {
		if i%1000000000 == 0 {
			fmt.Println("\n**************")
			frame = getNextFrame(frame)
			for row := 0; row < height; row++ {
				printRow(frame[row])
			}
			fmt.Println("\n**************")
		}
	}
}
