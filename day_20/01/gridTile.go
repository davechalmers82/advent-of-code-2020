package main

type GridTile struct {
	tile *TileDefinition
	flipX bool
	flipY bool
}

func reverse(bytesList []byte) []byte {
	newList := make([]byte, 0, len(bytesList))
	for i := len(bytesList)-1; i >= 0; i-- {
		newList = append(newList, bytesList[i])
	}
	return newList
}

func (g GridTile) Top() []byte {
	top := make([]byte, len(g.tile.data[0]))

	if !g.flipY {
		copy(top, g.tile.data[0])
	} else {
		copy(top, g.tile.data[len(g.tile.data)-1])
	}

	if g.flipX {
		top = reverse(top)
	}

	return top
}

func (g GridTile) Bottom() []byte {
	bottom := make([]byte, len(g.tile.data[0]))

	if g.flipY {
		copy(bottom, g.tile.data[0])
	} else {
		copy(bottom, g.tile.data[len(g.tile.data)-1])
	}

	if g.flipX {
		bottom = reverse(bottom)
	}

	return bottom
}

func (g GridTile) Left() []byte {
	left := make([]byte, len(g.tile.data))

	xIndex := 0
	if g.flipX {
		xIndex = len(g.tile.data[0])-1
	}

	for y := 0; y < len(g.tile.data); y++ {
		left[y] = g.tile.data[y][xIndex]
	}

	if g.flipY {
		left = reverse(left)
	}

	return left
}

func (g GridTile) Right() []byte {
	right := make([]byte, len(g.tile.data))

	xIndex := 0
	if !g.flipX {
		xIndex = len(g.tile.data[0])-1
	}

	for y := 0; y < len(g.tile.data); y++ {
		right[y] = g.tile.data[y][xIndex]
	}

	if g.flipY {
		right = reverse(right)
	}

	return right
}

func NewGridTile(def *TileDefinition, flipX bool, flipY bool) *GridTile {
	return &GridTile{
		tile: def,
		flipY: flipY,
		flipX: flipX,
	}
}

