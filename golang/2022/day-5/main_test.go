package main

import (
	"testing"

	"github.com/cassis163/advent-of-code/util"
)

func TestPartOne(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	got := PartOne(data)
	want := "CMZ"
	if got != want {
		t.Errorf("PartOne(data) = %s; want %s", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	got := PartTwo(data)
	want := "MCD"
	if got != want {
		t.Errorf("PartTwo(data) = %s; want %s", got, want)
	}
}
