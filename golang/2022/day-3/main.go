package main

import (
	"fmt"
	"strings"

	"github.com/cassis163/advent-of-code/util"
)

type Rucksack struct {
	compartmentA, compartmentB string
}

func main() {
	data, err := util.ReadFileAsString("./data.txt")
	if err != nil {
		panic(err)
	}

	rucksacks := getRucksacks(data)
	// Item types that are present in both compartments of a rucksack
	foundItemTypes := getItemsInRucksacks(rucksacks)
	prioritySum := getPrioritySum(foundItemTypes)

	fmt.Printf("Item types: %s\n", foundItemTypes)
	fmt.Printf("Priority sum: %d\n", prioritySum)
}

func getPrioritySum(itemTypes []string) int {
	var prioritySum int
	for _, itemType := range itemTypes {
		prioritySum += getPriorityOfItemType(itemType)
	}

	return prioritySum
}

func getItemsInRucksacks(rucksacks []Rucksack) []string {
	var foundItemTypes []string
	for _, rucksack := range rucksacks {
		for _, itemType := range rucksack.compartmentA {
			if strings.Contains(rucksack.compartmentB, string(itemType)) {
				foundItemTypes = append(foundItemTypes, string(itemType))
				break
			}
		}
	}

	return foundItemTypes
}

func getRucksacks(data string) []Rucksack {
	var rucksacks []Rucksack
	for _, line := range strings.Split(data, "\n") {
		compartmentA := line[:len(line)/2]
		compartmentB := line[len(line)/2:]

		rucksack := Rucksack{compartmentA, compartmentB}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func getPriorityOfItemType(itemType string) int {
	if itemType == strings.ToUpper(itemType) {
		// Uppercase item types A through Z have priorities 27 through 52.
		return int(itemType[0]) - 38
	}

	// Lowercase item types a through z have priorities 1 through 26.
	return int(itemType[0] - 96)
}
