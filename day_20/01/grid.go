package main

import (
	"bytes"
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

func (g *Grid) Width() int {
	return len(g.grid[0])
}

func (g *Grid) Height() int {
	return len(g.grid)
}

func (g *Grid) Clone() *Grid {
	// Clone the grid
	clonedGrid := NewGrid(g.Width(), g.Height(), g.available)
	copy(clonedGrid.grid, g.grid)
	for i, row := range g.grid {
		copy(clonedGrid.grid[i], row)
	}
	return clonedGrid
}

func (g *Grid) FindEmptyGridTile() (x int, y int) {
	for y, row := range g.grid {
		for x, gridTile := range row {
			if gridTile.tile == nil {
				return x, y
			}
		}
	}
	return -1, -1
}

func (g *Grid) AllPossibleGridTilesForCoord(x int, y int) (possible []*GridTile) {

	for _, def := range g.available {

		standard := NewGridTile(def, false, false)
		if g.CanBePlaced(x, y, standard) {
			possible = append(possible, standard)
		}

		flippedX := NewGridTile(def, true, false)
		if g.CanBePlaced(x, y, flippedX) {
			possible = append(possible, flippedX)
		}

		flippedY := NewGridTile(def, false, true)
		if g.CanBePlaced(x, y, flippedY) {
			possible = append(possible, flippedY)
		}

		flippedXY := NewGridTile(def, true, true)
		if g.CanBePlaced(x, y, flippedXY) {
			possible = append(possible, flippedXY)
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

	g.grid[y][x].tile = gridTile.tile
	g.grid[y][x].flipX = gridTile.flipX
	g.grid[y][x].flipY = gridTile.flipY

	// Remove from available
	index := g.indexOfAvailableDefinition(gridTile.tile)
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
	if north != nil && north.tile != nil && 0 != bytes.Compare(north.Bottom(), gridTile.Top()) {
		return false
	}

	south := g.at(x, y+1)
	if south != nil && south.tile != nil && 0 != bytes.Compare(south.Top(), gridTile.Bottom()) {
		return false
	}

	east := g.at(x+1, y)
	if east != nil && east.tile != nil && 0 != bytes.Compare(east.Left(), gridTile.Right()) {
		return false
	}

	west := g.at(x-1, y)
	if west != nil && west.tile != nil && 0 != bytes.Compare(west.Right(), gridTile.Left()) {
		return false
	}

	return true
}
