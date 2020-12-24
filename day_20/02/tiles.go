package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TileDefinition struct {
	id int
	data [][]byte
}

func (t TileDefinition) Print()  {
	fmt.Println("Tile: " + strconv.Itoa(t.id))
	for i, row := range t.data {
		fmt.Println(i, string(row))
	}
}

func LoadTilesFromFile(path string) (tiles []*TileDefinition, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tile *TileDefinition = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()

		if strings.HasPrefix(lineText,"Tile") {
			tile = &TileDefinition{}
			_, err := fmt.Sscanf(lineText, "Tile %d:", &tile.id)
			if err != nil {
				return nil, err
			}
		} else if tile != nil {
			if len(strings.TrimSpace(lineText)) == 0 {
				tiles = append(tiles, tile)
				tile = nil
			} else {
				// Add the row to the tile
				tile.data = append(tile.data, []byte(lineText))
			}
		}
	}

	if tile != nil {
		tiles = append(tiles, tile)
		tile = nil
	}

	return tiles, scanner.Err()
}
