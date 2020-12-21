package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadFileNumberArray(path string) (numbers []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, value)
	}
	return numbers, scanner.Err()
}

func main() {

	// Load the file
	numbers, err := loadFileNumberArray("../input.txt")
	if err != nil {
		fmt.Printf("File loading failed!", err)
		os.Exit(1)
	}

	for idx, first := range numbers {
		for idx2, second := range numbers[idx+1:] {
			for _, third := range numbers[idx2+1:] {

				sum := first + second + third

				if sum == 2020 {
					fmt.Printf("a=%d, b=%d c=%d sum=%d res=%d", first, second, third, sum, first*second*third)
					return
				}
			}
		}
	}

	fmt.Printf("No result found")
}
