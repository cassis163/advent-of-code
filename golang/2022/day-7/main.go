package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/cassis163/advent-of-code/util"
)

type File struct {
	SizeInBytes int
	Name        string
	parent      *Directory
}

type Directory struct {
	Files       []File
	Directories []Directory
	Name        string
	parent      *Directory
}

func (d Directory) Size() int {
	total := 0
	for _, file := range d.Files {
		total += file.SizeInBytes
	}
	for _, dir := range d.Directories {
		total += dir.Size()
	}
	return total
}

func (d Directory) GetSubDirectory(name string) (*Directory, error) {
	for _, dir := range d.Directories {
		if dir.Name == name {
			return &dir, nil
		}
	}
	return nil, fmt.Errorf("directory %s not found in %s", name, d.Name)
}

func (d Directory) Print() {
	var prettyPrintDirectory func(dir Directory, indent string)
	prettyPrintDirectory = func(dir Directory, indent string) {
		fmt.Printf("%sDirectory: %s\n", indent, dir.Name)

		for _, file := range dir.Files {
			fmt.Printf("%s- File: %s (%d bytes)\n", indent, file.Name, file.SizeInBytes)
		}

		for _, subdir := range dir.Directories {
			prettyPrintDirectory(subdir, indent+"  ")
		}
	}

	prettyPrintDirectory(d, "")
}

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

	directoriesToBeRemoved := getDirectoriesToBeRemoved(rootDirectory)
	total := 0
	for _, dir := range directoriesToBeRemoved {
		total += dir.Size()
	}

	fmt.Printf("Part one: %d\n", total)

	return total
}

func PartTwo(data string) int {
	return 0
}

func deserialize(data string) (Directory, error) {
	rootDirectory := Directory{}
	currentDir := &rootDirectory
	for _, line := range strings.Split(data, "\n") {
		err := deserializeLine(line, currentDir)
		if err != nil {
			return Directory{}, err
		}
	}

	return rootDirectory, nil
}

func deserializeLine(line string, currentDir *Directory) error {
	switch {
	case line == "$ cd /" || line == "$ ls" || line == "":
		return nil
	case line == "$ cd ..":
		fmt.Printf("currentDir: %s\n", currentDir.Name)
		*currentDir = *currentDir.parent
	case strings.HasPrefix(line, "$ cd "):
		dirName := strings.TrimPrefix(line, "$ cd ")
		subDir, err := currentDir.GetSubDirectory(dirName)
		if err != nil {
			return err
		}
		*currentDir = *subDir
	case strings.HasPrefix(line, "dir "):
		dirName := strings.TrimPrefix(line, "dir ")
		newDirectory := Directory{Name: dirName, parent: currentDir}
		currentDir.Directories = append(currentDir.Directories, newDirectory)
	case unicode.IsDigit(rune(line[0])):
		parts := strings.Split(line, " ")
		sizeInBytes, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		fileName := parts[1]
		newFile := File{SizeInBytes: sizeInBytes, Name: fileName, parent: currentDir}
		currentDir.Files = append(currentDir.Files, newFile)
	default:
		return fmt.Errorf("invalid line: %s", line)
	}

	return nil
}
