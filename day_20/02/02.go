package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"math"
	"os"
)

var (
	all = kingpin.Flag("all", "Find all solutions.").Short('a').Bool()
)

func main() {
	kingpin.Parse()

	// Load the file
	tiles, err := LoadTilesFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	size := int(math.Sqrt(float64(len(tiles))))
	completedGrids := completeGrid(NewGrid(size, size, tiles), *all)

	for i, solution := range completedGrids {
		fmt.Printf("Solution: %d/%d\n", i + 1, len(completedGrids))
		solution.Print()

		fullMap := solution.GenerateSingleBuffer(0, true)
		printBuffer(&fullMap)
	}
}

func printBuffer(buffer *[][]byte) {
	for _, row := range *buffer {
		fmt.Println(string(row))
	}
}

func completeGrid(grid *Grid, findAll bool) (completeGrids []*Grid) {
	x, y := grid.FindEmptyGridTile()
	if x >= 0 && y >= 0 {
		possibleGridTiles := grid.AllPossibleGridTilesForCoord(x, y)

		for _, gridTile := range possibleGridTiles {

			newGrid := grid.Clone()
			newGrid.PlaceTile(x, y, gridTile)

			completedGrid := completeGrid(newGrid, findAll)
			if len(completedGrid) > 0 {
				completeGrids = append(completeGrids, completedGrid...)

				if !findAll {
					return completeGrids
				}
			}
		}
	} else {
		completeGrids = append(completeGrids, grid)
	}

	return completeGrids
}