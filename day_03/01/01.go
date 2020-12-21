package main

import (
	"bufio"
	"fmt"
	"os"
)

type Map struct {
	grid [][]byte
}

func (m Map) Print()  {
	for _, row := range m.grid {
		fmt.Println(string(row))
	}
}

func (m Map) valueAt(x int, y int) byte {
	// Handle infinite scroll in X
	cycleX := x % len(m.grid[y])
	return m.grid[y][cycleX]
}

func (m Map) GenerateRoute(right int, down int) []byte {
	var route []byte

	x := 0
	for y := 0; y < len(m.grid); y += down {
		route = append(route, m.valueAt(x, y))
		x += right
	}

	return route
}

func loadMapFromFile(path string) (grid *Map, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	newMap := Map{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newMap.grid = append(newMap.grid, []byte(scanner.Text()))
	}
	return &newMap, scanner.Err()
}

func countTrees(route []byte) int {
	treeCount := 0
	for _, value := range route {
		if value == '#' {
			treeCount++
		}
	}
	return treeCount
}

func main() {
	// Load the file
	newMap, err := loadMapFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	route := newMap.GenerateRoute(3, 1)
	treeCount := countTrees(route)

	fmt.Printf("Tree Count = %d\n", treeCount)
}