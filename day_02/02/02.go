package main

import (
	"bufio"
	"fmt"
	"os"
)

type Entry struct {
	Min int
	Max int
	Char rune
	Password string
}

func (e Entry) Print() {
	fmt.Printf("%d-%d %c: %s\n", e.Min, e.Max, e.Char, e.Password)
}

// Must contain only 1 in the min/max position
func (e Entry) IsValid() bool {

	runes := []rune(e.Password)

	matches := 0

	if runes[e.Min-1] == e.Char {
		matches++
	}

	if runes[e.Max-1] == e.Char {
		matches++
	}

	return matches == 1
}

func LoadEntriesFromFile(path string) (entries []Entry, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		newEntry := Entry{}
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &newEntry.Min, &newEntry.Max, &newEntry.Char, &newEntry.Password)
		if err != nil {
			return  nil, err
		}
		entries = append(entries, newEntry)
	}
	return entries, scanner.Err()
}

func main() {
	// Load the file
	entries, err := LoadEntriesFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	validCount := 0
	for _, entry := range entries {
		if entry.IsValid() {
			validCount++
			entry.Print()
		}
	}

	fmt.Printf("Valid Count: %d", validCount)
}
