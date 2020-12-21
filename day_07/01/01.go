package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type BagContents struct {
	bags map[string]int
}

func (c BagContents) eventuallyContains(searchId string, definitions map[string]*BagContents) bool {
	for id := range c.bags {
		if id == searchId {
			return true
		}

		if definitions[id] != nil && definitions[id].eventuallyContains(searchId, definitions) {
			return true
		}
	}
	return false
}

func parseBagContents(str string) *BagContents {
	bags := strings.Split(str, ",")

	contents := BagContents{
		bags: make(map[string]int),
	}

	for _, bag := range bags {
		var count int
		var name1, name2 string
		_, err := fmt.Sscanf(bag, "%d %s %s", &count, &name1, &name2)
		if err != nil {
			return  nil
		}

		contents.bags[name1 + " " + name2 ] = count
	}

	return &contents
}

func loadBagsFromFile(path string) (bags map[string]*BagContents, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var re = regexp.MustCompile(`(?P<id>[a-z]+ [a-z]+) bags contain (?:(?P<contents>[0-9]+ .*).)?`)
	idIndex := re.SubexpIndex("id")
	contentsIndex := re.SubexpIndex("contents")

	bags = make(map[string]*BagContents)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		bags[match[idIndex]] = parseBagContents(match[contentsIndex])
	}
	return bags, scanner.Err()
}

func main() {
	// Load the file
	bagDefinitions, err := loadBagsFromFile("../input.txt")
	if err != nil {
		fmt.Print("File loading failed!", err)
		os.Exit(1)
	}

	const searchBagId = "shiny gold"

	count := 0

	for id, bag := range bagDefinitions {
		if id != searchBagId && bag != nil {
			if bag.eventuallyContains(searchBagId, bagDefinitions) {
				count++
			}
		}
	}
	fmt.Printf("%d bags contain '%s' bag!", count, searchBagId)
}
