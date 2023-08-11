package main

import (
	"fmt"
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

func (d *Directory) AddFile(file File) {
	file.parent = d
	d.Files = append(d.Files, file)
}

func (d *Directory) AddDirectory(dir Directory) {
	dir.parent = d
	d.Directories = append(d.Directories, dir)
}

func (d *Directory) Size() int {
	total := 0
	for _, file := range d.Files {
		total += file.SizeInBytes
	}
	for _, dir := range d.Directories {
		total += dir.Size()
	}
	return total
}

func (d *Directory) GetSubDirectory(name string) (*Directory, error) {
	for _, dir := range d.Directories {
		if dir.Name == name {
			return &dir, nil
		}
	}
	fmt.Printf("directory %v\n", d)
	return nil, fmt.Errorf("directory %s not found in %s", name, d.Name)
}

func (d *Directory) Print() {
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

	prettyPrintDirectory(*d, "")
}
