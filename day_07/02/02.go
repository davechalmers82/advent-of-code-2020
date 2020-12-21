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

func (c BagContents) countAllBags(definitions map[string]*BagContents) int {
	total := 1

	for id, count := range c.bags {
		if definitions[id] != nil {
			total += count * definitions[id].countAllBags(definitions)
		} else {
			total += count
		}
	}
	return total
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

	// NOTE: Need to take the one off for this bag!
	const searchBagId = "shiny gold"
	count := bagDefinitions[searchBagId].countAllBags(bagDefinitions) - 1
	fmt.Printf("%d bags contain '%s' bag!", count, searchBagId)
}
