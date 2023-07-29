package main

import (
	"testing"

	"github.com/cassis163/advent-of-code/util"
)

func TestGetScoreForPartOne(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	turns := getTurns(data)

	want := 15
	got := getScoreForPartOne(turns)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestGetScoreForPartTwo(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	turns := getTurns(data)

	want := 12
	got := getScoreForPartTwo(turns)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
