package main

type GridTile struct {
	tileDef *TileDefinition
	modifiedData [][]byte
	flipX bool
	flipY bool
	rotate bool
}

func reverse(bytesList []byte) []byte {
	newList := make([]byte, 0, len(bytesList))
	for i := len(bytesList)-1; i >= 0; i-- {
		newList = append(newList, bytesList[i])
	}
	return newList
}

func flipInX(data [][]byte) [][]byte {
	flipped := make([][]byte, len(data))
	for i := range data {
		flipped[i] = reverse(data[i])
	}
	return flipped
}

func flipInY(data [][]byte) [][]byte {
	flipped := make([][]byte, len(data))

	for y := 0; y < len(data); y++ {
		srcIndex := len(data) - 1 - y

		flipped[y] = make([]byte, 0, len(data[y]))
		flipped[y] = append(flipped[y], data[srcIndex]...)
	}
	return flipped
}

func rotateCW(data [][]byte) [][]byte {
	flipped := make([][]byte, len(data))

	for y := 0; y < len(data); y++ {
		flipped[y] = make([]byte, 0, len(data[y]))

		for i := 0; i < len(data); i++ {
			flipped[y] = append(flipped[y], data[len(data) - 1 - i][y])
		}

	}
	return flipped
}


func (g GridTile) Top() []byte {
	return g.modifiedData[0]
}

func (g GridTile) Bottom() []byte {
	return g.modifiedData[len(g.modifiedData)-1]
}

func (g GridTile) Left() []byte {
	left := make([]byte, len(g.modifiedData))

	for y := 0; y < len(g.modifiedData); y++ {
		left[y] = g.modifiedData[y][0]
	}
	return left
}

func (g GridTile) Right() []byte {
	right := make([]byte, len(g.modifiedData))

	xIndex := len(g.modifiedData[0])-1

	for y := 0; y < len(g.modifiedData); y++ {
		right[y] = g.modifiedData[y][xIndex]
	}

	return right
}

func NewGridTile(def *TileDefinition, flipX bool, flipY bool, rotate bool) *GridTile {

	gridTile := &GridTile{
		tileDef: def,
		flipY: flipY,
		flipX: flipX,
		rotate: rotate,
	}

	gridTile.modifiedData = make([][]byte, len(def.data))
	for i := range def.data {
		gridTile.modifiedData[i] = make([]byte, len(def.data[i]))
		copy(gridTile.modifiedData[i], def.data[i])
	}

	if rotate {
		gridTile.modifiedData = rotateCW(gridTile.modifiedData)
	}

	if flipX {
		gridTile.modifiedData = flipInX(gridTile.modifiedData)
	}

	if flipY {
		gridTile.modifiedData = flipInY(gridTile.modifiedData)
	}

	return gridTile
}

