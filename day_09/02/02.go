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

func findSet(startIndex int, numbers []int) []int {
	const target = 18272118

	sum := 0
	for i := startIndex; i < len(numbers); i++ {
		sum += numbers[i]
		if sum == target {
			return numbers[startIndex:i]
		} else if sum > target {
			break
		}
	}
	return nil
}

func minMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	// Load the file
	numbers, err := loadNumbersFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	for i := 0; i < len(numbers); i++ {
		set := findSet(i, numbers)
		if set != nil {
			min, max := minMax(set)
			fmt.Printf("%d+%d=%d\n", min, max, min+max)
			break
		}
	}
}
