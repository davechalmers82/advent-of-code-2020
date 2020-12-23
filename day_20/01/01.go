package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// Load the file
	tiles, err := LoadTilesFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	//for i, tile := range tiles {
	//	tile.Print()
//		fmt.Println(i)
//	}

	size := int(math.Sqrt(float64(len(tiles))))

	fmt.Println("======================================")
	fmt.Println(size)
	fmt.Println("======================================")

	completedGrids := completeGrid(NewGrid(size, size, tiles))

	fmt.Println(len(completedGrids))
}

func completeGrid(grid *Grid) (completeGrids []*Grid) {

	x, y := grid.FindEmptyGridTile()
	if x >= 0 && y >= 0 {
		possibleGridTiles := grid.AllPossibleGridTilesForCoord(x, y)

		for _, gridTile := range possibleGridTiles {

			newGrid := grid.Clone()
			newGrid.PlaceTile(x, y, gridTile)

			completedGrid := completeGrid(newGrid)
			if len(completedGrid) > 0 {
				completeGrids = append(completeGrids, completedGrid...)
			}
		}
	} else {
		completeGrids = append(completeGrids, grid)
	}

	return completeGrids
}