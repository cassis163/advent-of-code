package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/cassis163/advent-of-code/util"
)

func main() {
	data, err := util.ReadFileAsString("./data.txt")
	if err != nil {
		panic(err)
	}

	PartOne(data)
	PartTwo(data)
}

func PartOne(data string) string {
	result, err := getCharacterIndexBeforeFirstMarker(data, 4)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part one: %d\n", result)
	return strconv.Itoa(result)
}

func PartTwo(data string) string {
	result, err := getCharacterIndexBeforeFirstMarker(data, 14)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part two: %d\n", result)
	return strconv.Itoa(result)
}

func getCharacterIndexBeforeFirstMarker(data string, distinctCharactersAmount int) (int, error) {
	for startIndex := 0; startIndex < len(data)-1; startIndex++ {
		endIndex := startIndex + distinctCharactersAmount
		characterSequence := data[startIndex:endIndex]
		if doesPrependMarker(characterSequence) {
			return endIndex, nil
		}
	}

	return 0, errors.New("no marker found")
}

func doesPrependMarker(characterSequence string) bool {
	for _, char := range characterSequence {
		if doesStringContainCharacterTwice(characterSequence, char) {
			return false
		}
	}

	return true
}

func doesStringContainCharacterTwice(characterSequence string, character rune) bool {
	count := 0
	for _, char := range characterSequence {
		if char == character {
			count++
			if count == 2 {
				return true
			}
		}
	}

	return false
}
