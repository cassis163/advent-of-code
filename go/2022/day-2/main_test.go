package main

import (
	"testing"

	"github.com/cassis163/advent-of-code/util"
)

func TestGetScore(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	turns := getTurns(data)

	want := 15
	got := getScore(turns)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
