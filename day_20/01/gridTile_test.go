package main

import (
	"bytes"
	"testing"
)

func createTestTileDefinition() *TileDefinition {
	return &TileDefinition{
		id: 666,
		data: [][]byte{
			{ '.', '.', '#', '#', '.', '#', '.', '.', '#', '.' },
			{ '#', '#', '.', '.', '#', '.', '.', '.', '.', '.' },
			{ '#', '.', '.', '.', '#', '#', '.', '.', '#', '.' },
			{ '#', '#', '#', '#', '.', '#', '.', '.', '.', '#' },
			{ '#', '#', '.', '#', '#', '.', '#', '#', '#', '.' },
			{ '#', '#', '.', '.', '.', '#', '.', '#', '#', '#' },
			{ '.', '#', '.', '#', '.', '#', '.', '.', '#', '#' },
			{ '.', '.', '#', '.', '.', '.', '.', '#', '.', '.' },
			{ '#', '#', '#', '.', '.', '.', '#', '.', '#', '.' },
			{ '.', '.', '#', '#', '#', '.', '.', '#', '#', '#' },
		},
	}
}

func TestNewGridTile(t *testing.T) {
	newTile := NewGridTile(createTestTileDefinition(), true, false)
	if newTile == nil {
		t.Error("expected a new GridTile got 'nil'")
	}
}

func TestUniqueGridTile(t *testing.T) {
	newTileA := NewGridTile(createTestTileDefinition(), true, false)
	newTileB := NewGridTile(createTestTileDefinition(), true, false)

	if newTileA == newTileB {
		t.Error("expected a unique GridTiles")
	}
}

func TestGridTileTop(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, false)

	top := tile.Top()

	if 0 != bytes.Compare(top, []byte{ '.', '.', '#', '#', '.', '#', '.', '.', '#', '.' }) {
		t.Error("result should match")
	}

	if &(tile.tile.data[0]) == &top {
		t.Error("pointers should not match!")
	}
}

func TestGridTileTopFlipY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, true)
	if 0 != bytes.Compare(tile.Top(), []byte{ '.', '.', '#', '#', '#', '.', '.', '#', '#', '#' }) {
		t.Error("result should match")
	}
}

func TestGridTileTopFlipX(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, false)
	if 0 != bytes.Compare(tile.Top(), []byte{ '.', '#', '.', '.', '#', '.', '#', '#', '.', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileTopFlipXY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, true)
	if 0 != bytes.Compare(tile.Top(), []byte{ '#', '#', '#', '.', '.', '#', '#', '#', '.', '.' }) {
		t.Error("result should match")
	}
}


func TestGridTileBottom(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, false)

	bottom := tile.Bottom()

	if 0 != bytes.Compare(bottom, []byte{ '.', '.', '#', '#', '#', '.', '.', '#', '#', '#' }) {
		t.Error("result should match")
	}

	if &(tile.tile.data[0]) == &bottom {
		t.Error("pointers should not match!")
	}
}

func TestGridTileBottomFlipY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, true)
	if 0 != bytes.Compare(tile.Bottom(), []byte{ '.', '.', '#', '#', '.', '#', '.', '.', '#', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileBottomFlipX(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, false)
	if 0 != bytes.Compare(tile.Bottom(), []byte{ '#', '#', '#', '.', '.', '#', '#', '#', '.', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileBottomFlipXY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, true)
	if 0 != bytes.Compare(tile.Bottom(), []byte{ '.', '#', '.', '.', '#', '.', '#', '#', '.', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileLeft(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, false)

	if 0 != bytes.Compare(tile.Left(), []byte{ '.', '#', '#', '#', '#', '#', '.', '.', '#', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileLeftFlipY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, true)
	if 0 != bytes.Compare(tile.Left(), []byte{ '.', '#', '.', '.', '#', '#', '#', '#', '#', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileLeftFlipX(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, false)
	if 0 != bytes.Compare(tile.Left(), []byte{ '.', '.', '.', '#', '.', '#', '#', '.', '.', '#' }) {
		t.Error("result should match")
	}
}

func TestGridTileLeftFlipXY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, true)
	if 0 != bytes.Compare(tile.Left(), []byte{ '#', '.', '.', '#', '#', '.', '#', '.', '.', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileRight(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, false)

	if 0 != bytes.Compare(tile.Right(), []byte{ '.', '.', '.', '#', '.', '#', '#', '.', '.', '#' }) {
		t.Error("result should match")
	}
}

func TestGridTileRightFlipY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), false, true)
	if 0 != bytes.Compare(tile.Right(), []byte{ '#', '.', '.', '#', '#', '.', '#', '.', '.', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileRightFlipX(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, false)
	if 0 != bytes.Compare(tile.Right(), []byte{ '.', '#', '#', '#', '#', '#', '.', '.', '#', '.' }) {
		t.Error("result should match")
	}
}

func TestGridTileRightFlipXY(t *testing.T) {
	tile := NewGridTile(createTestTileDefinition(), true, true)
	if 0 != bytes.Compare(tile.Right(), []byte{ '.', '#', '.', '.', '#', '#', '#', '#', '#', '.' }) {
		t.Error("result should match")
	}
}
