package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadNumbersFromFile(path string) (numbers []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var value int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &value)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, value)
	}
	return numbers, scanner.Err()
}

func isValid(value int, numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i+1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == value {
				return true
			}
		}
	}
	return false
}

func main() {
	// Load the file
	numbers, err := loadNumbersFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	const size = 25

	for i := size; i < len(numbers); i++ {
		if !isValid(numbers[i], numbers[i-size:i]) {
			fmt.Printf("(%d) Value: %d\n", i, numbers[i])
			break
		}
	}
}
