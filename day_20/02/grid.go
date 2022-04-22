package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Grid struct {
	grid [][]GridTile
	available []*TileDefinition
}

func NewGrid(width int, height int, tiles []*TileDefinition) *Grid {
	// Clone the grid
	newGrid := &Grid{}

	newGrid.available = make([]*TileDefinition, len(tiles))
	copy(newGrid.available, tiles)

	newGrid.grid = make([][]GridTile, height)
	for i := range newGrid.grid {
		newGrid.grid[i] = make([]GridTile, width)
	}

	return newGrid
}

func (g *Grid) Print() {
	fmt.Println("-----------------------------------------------------------------------------------")
	for _, row := range g.grid {
		fmt.Print("| ")
		for _, item := range row {
			if item.tileDef != nil {
				fmt.Print(strconv.Itoa(item.tileDef.id) + " | ")
			} else {
				fmt.Print("NONE | ")
			}
		}
		fmt.Println()
	}
	fmt.Println("-----------------------------------------------------------------------------------")
}

func (g *Grid) GenerateSingleBuffer(spacers int, removeBorder bool) [][]byte {
	tileSizeX := len(g.grid[0][0].tileDef.data[0])
	tileSizeY := len(g.grid[0][0].tileDef.data)

	border := 0

	if removeBorder {
		tileSizeX -= 2
		tileSizeY -= 2
		border = 1
	}

	spacesX := (g.Width() - 1) * spacers
	spacesY := (g.Height() - 1) * spacers

	buffer := make([][]byte, (g.Height() * tileSizeY) + spacesX)
	for y := range buffer {
		buffer[y] = make([]byte, (g.Width() * tileSizeX) + spacesY)
	}

	if spacers > 0 {
		defaultValueIntoBuffer(&buffer, ' ')
	}

	for tileY, row := range g.grid {
		for tileX, tile := range row {
			if tile.modifiedData != nil {
				startX := (tileX * tileSizeX) + (spacers * tileX)
				startY := (tileY * tileSizeY) + (spacers * tileY)
				copyBufferIntoBuffer(&buffer, startX, startY, tile.modifiedData, border)
			}
		}
	}
	return buffer
}

func defaultValueIntoBuffer(buffer *[][]byte, value byte) {
	for y, row := range *buffer {
		for x := range row {
			(*buffer)[y][x] = value
		}
	}
}

func copyBufferIntoBuffer(buffer *[][]byte, destX int, destY int, data [][]byte, border int) {
	for y := border; y < len(data) - border; y++ {
		for x := border; x < len(data[y]) - border; x++ {
			(*buffer)[destY+y-border][destX+x-border] = data[y][x]
		}
	}
}

func (g *Grid) Width() int {
	return len(g.grid[0])
}

func (g *Grid) Height() int {
	return len(g.grid)
}

func (g *Grid) Clone() *Grid {
	// Clone the grid
	clonedGrid := NewGrid(g.Width(), g.Height(), g.available)
	for y, row := range g.grid {
		for x, item := range row {
			clonedGrid.grid[y][x].tileDef = item.tileDef
			clonedGrid.grid[y][x].modifiedData = item.modifiedData
			clonedGrid.grid[y][x].flipY = item.flipY
			clonedGrid.grid[y][x].flipX = item.flipX
			clonedGrid.grid[y][x].rotate = item.rotate
		}
	}
	return clonedGrid
}

func (g *Grid) FindEmptyGridTile() (x int, y int) {
	for y, row := range g.grid {
		for x, gridTile := range row {
			if gridTile.tileDef == nil {
				return x, y
			}
		}
	}
	return -1, -1
}

func (g *Grid) AllPossibleGridTilesForCoord(x int, y int) (possible []*GridTile) {
	for _, def := range g.available {
		// Do all these twice once not rotated and once with
		for i := 0; i < 2; i++ {
			standard := NewGridTile(def, false, false, i == 1)
			if g.CanBePlaced(x, y, standard) {
				possible = append(possible, standard)
			}

			flippedX := NewGridTile(def, true, false, i == 1)
			if g.CanBePlaced(x, y, flippedX) {
				possible = append(possible, flippedX)
			}

			flippedY := NewGridTile(def, false, true, i == 1)
			if g.CanBePlaced(x, y, flippedY) {
				possible = append(possible, flippedY)
			}

			flippedXY := NewGridTile(def, true, true, i == 1)
			if g.CanBePlaced(x, y, flippedXY) {
				possible = append(possible, flippedXY)
			}
		}
	}

	return possible
}

func (g *Grid) indexOfAvailableDefinition(def *TileDefinition) (index int) {
	for i, item := range g.available {
		if item == def {
			return i
		}
	}

	return -1
}

func (g *Grid) PlaceTile(x int, y int, gridTile *GridTile) {

	g.grid[y][x].tileDef = gridTile.tileDef
	g.grid[y][x].modifiedData = gridTile.modifiedData
	g.grid[y][x].flipX = gridTile.flipX
	g.grid[y][x].flipY = gridTile.flipY
	g.grid[y][x].rotate = gridTile.rotate

	// Remove from available
	index := g.indexOfAvailableDefinition(gridTile.tileDef)
	if index >= 0 {
		g.available = append(g.available [:index], g.available [index+1:]...)
	}
}

func (g *Grid) at(x int, y int) (gridTile *GridTile) {
	if x >= 0 && y >= 0 && x < g.Width() && y < g.Height() {
		return &g.grid[y][x]
	}
	return nil
}

func (g *Grid) CanBePlaced(x int, y int, gridTile *GridTile) bool {

	north := g.at(x, y-1)
	if north != nil && north.tileDef != nil && 0 != bytes.Compare(north.Bottom(), gridTile.Top()) {
		return false
	}

	south := g.at(x, y+1)
	if south != nil && south.tileDef != nil && 0 != bytes.Compare(south.Top(), gridTile.Bottom()) {
		return false
	}

	east := g.at(x+1, y)
	if east != nil && east.tileDef != nil && 0 != bytes.Compare(east.Left(), gridTile.Right()) {
		return false
	}

	west := g.at(x-1, y)
	if west != nil && west.tileDef != nil && 0 != bytes.Compare(west.Right(), gridTile.Left()) {
		return false
	}

	return true
}
