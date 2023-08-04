package main

import (
	"regexp"
	"strconv"

	"github.com/cassis163/advent-of-code/util"
)

func main() {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	PartOne(data)
	PartTwo(data)
}

func PartOne(data string) string {
	return strconv.Itoa(getTotalSizeOfDirectories(data))
}

func PartTwo(data string) string {
	return ""
}

func getTotalSizeOfDirectories(data string) int {
	re := regexp.MustCompile("[0-9]+")
	fileSizes := re.FindAllString(data, -1)

	totalSize := 0
	for _, fileSize := range fileSizes {
		size, err := strconv.Atoi(fileSize)
		if err != nil {
			panic(err)
		}

		totalSize += size
	}

	return totalSize
}
