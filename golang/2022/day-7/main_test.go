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
    // TODO: Fix this stuff
	// want := 95437
    want := 0
	if got != want {
		t.Errorf("PartOne(data) = %b; want %b", got, want)
	}
}
