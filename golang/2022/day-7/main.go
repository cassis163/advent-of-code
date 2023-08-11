package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

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

func PartOne(data string) int {
	rootDirectory, err := deserialize(data)
	if err != nil {
		panic(err)
	}

	rootDirectory.Print()

	var getDirectoriesToBeRemoved func(dir Directory) []Directory
	getDirectoriesToBeRemoved = func(dir Directory) []Directory {
		if dir.Size() <= 100000 {
			return []Directory{dir}
		}
		result := []Directory{}
		for _, subDir := range dir.Directories {
			result = append(result, getDirectoriesToBeRemoved(subDir)...)
		}
		return result
	}

	directoriesToBeRemoved := getDirectoriesToBeRemoved(*rootDirectory)
	total := 0
	for _, dir := range directoriesToBeRemoved {
		total += dir.Size()
	}

	fmt.Printf("\nPart one: %d\n", total)

	return total
}

func PartTwo(data string) int {
	return 0
}

func deserialize(data string) (*Directory, error) {
	rootDirectory := Directory{Name: "root"}
	currentDir := &rootDirectory
	for _, line := range strings.Split(data, "\n") {
		err := deserializeLine(line, &currentDir)
		if err != nil {
			return &Directory{}, err
		}
	}

	fmt.Println()

	return &rootDirectory, nil
}

func deserializeLine(line string, currentDir **Directory) error {
	switch {
	case line == "$ cd /" || line == "$ ls" || line == "":
		return nil
	case line == "$ cd ..":
		*currentDir = (*currentDir).parent
	case strings.HasPrefix(line, "$ cd "):
		dirName := strings.TrimPrefix(line, "$ cd ")
		subDir, err := (*currentDir).GetSubDirectory(dirName)
		if err != nil {
			return err
		}
		*currentDir = subDir
	case strings.HasPrefix(line, "dir "):
		dirName := strings.TrimPrefix(line, "dir ")
		newDirectory := Directory{Name: dirName}
		(*currentDir).AddDirectory(newDirectory)
		fmt.Printf("Add %s to %s\n", dirName, (*currentDir).Name)
	case unicode.IsDigit(rune(line[0])):
		parts := strings.Split(line, " ")
		sizeInBytes, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		fileName := parts[1]
		newFile := File{SizeInBytes: sizeInBytes, Name: fileName}
		(*currentDir).AddFile(newFile)
		fmt.Printf("Add %s to %s\n", fileName, (*currentDir).Name)
	default:
		return fmt.Errorf("invalid line: %s", line)
	}

	return nil
}
