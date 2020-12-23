package main

import "testing"

func TestNewGrid(t *testing.T) {
	const height = 10
	const width = 12

	grid := NewGrid(width, height, nil)
	if grid == nil {
		t.Error("expected a new GridTile got 'nil'")
	} else {
		if grid.Height() != height {
			t.Error("expected a grid height not found")
		}
		if grid.Width() != width {
			t.Error("expected a grid width not found")
		}
	}
}

func TestGridClone(t *testing.T) {
	const height = 10
	const width = 12

	grid := NewGrid(width, height, nil)
	clonedGrid := grid.Clone()

	if grid.Height() != clonedGrid.Height() {
		t.Error("expected the same height")
	}
	if grid.Width() != clonedGrid.Width() {
		t.Error("expected the same width")
	}
	if grid == clonedGrid {
		t.Error("expected a different pointers for cloned grid")
	}
	if &grid.grid == &clonedGrid.grid {
		t.Error("expected a different data pointers for cloned grid")
	}
	if &grid.available == &clonedGrid.available {
		t.Error("expected a different available pointers for cloned grid")
	}
}

func TestGridPlaceTile(t *testing.T) {
	const height = 10
	const width = 12

	tileDef := createTestTileDefinition()

	grid := NewGrid(width, height, []*TileDefinition{ tileDef })
	gridTile := NewGridTile(tileDef, false, true, false)

	grid.PlaceTile(0, 0, gridTile)

	placedTile := grid.at(0, 0)

	if placedTile.tileDef != gridTile.tileDef {
		t.Error("expected same tile at 0, 0")
	}
	if placedTile.flipY != gridTile.flipY {
		t.Error("expected same tile flip at 0, 0")
	}
	if placedTile.flipX != gridTile.flipX {
		t.Error("expected same tile flip at 0, 0")
	}
	if len(grid.available) != 0 {
		t.Error("expected tile def removed from available")
	}
}

func TestGridPlaceTileCloned(t *testing.T) {
	const height = 10
	const width = 12

	tileDef := createTestTileDefinition()

	grid := NewGrid(width, height, []*TileDefinition{ tileDef })
	clonedGrid := grid.Clone()

	gridTile := NewGridTile(tileDef, false, true, false)

	// Place in cloned
	clonedGrid.PlaceTile(0, 0, gridTile)
	clonedTile := clonedGrid.at(0, 0)

	if clonedTile.tileDef != gridTile.tileDef {
		t.Error("expected same tile at 0, 0")
	}
	if clonedTile.flipY != gridTile.flipY {
		t.Error("expected same tile flip at 0, 0")
	}
	if clonedTile.flipX != gridTile.flipX {
		t.Error("expected same tile flip at 0, 0")
	}
	if len(clonedGrid.available) != 0 {
		t.Error("expected tile def removed from available")
	}

	// Check original
	originalTile := grid.at(0, 0)
	if originalTile.tileDef != nil {
		t.Error("should be unchanged")
	}
	if originalTile.flipY != false {
		t.Error("should be unchanged")
	}
	if originalTile.flipX != false {
		t.Error("should be unchanged")
	}
	if len(grid.available) != 1 {
		t.Error("should be unchanged")
	}
}

func TestGridFindEmptyGridTileEmpty(t *testing.T) {
	const height = 3
	const width = 3

	tileDef := createTestTileDefinition()

	grid := NewGrid(width, height, []*TileDefinition{ tileDef })

	x, y := grid.FindEmptyGridTile()

	if x != 0 || y != 0 {
		t.Error("when empty should find x=0, y=0")
	}
}

func TestGridFindEmptyGridTileAfterPlaced(t *testing.T) {
	const height = 3
	const width = 3

	tileDef := createTestTileDefinition()

	grid := NewGrid(width, height, []*TileDefinition{ tileDef })

	gridTile := NewGridTile(tileDef, false, true, false)
	grid.PlaceTile(0, 0, gridTile)

	x, y := grid.FindEmptyGridTile()
	if x != 1 || y != 0 {
		t.Error("when empty should find x=1, y=0")
	}
}

func TestGridFindEmptyGridTileFull(t *testing.T) {
	const height = 3
	const width = 3

	tileDef := createTestTileDefinition()

	grid := NewGrid(width, height, []*TileDefinition{ tileDef })

	gridTile := NewGridTile(tileDef, false, true, false)

	for i:= 0; i < height; i++ {
		for j:= 0; j < width; j++ {
			grid.PlaceTile(j, i, gridTile)
		}
	}

	x, y := grid.FindEmptyGridTile()
	if x != -1 || y != -1 {
		t.Error("should return -1, -1 when full")
	}
}


