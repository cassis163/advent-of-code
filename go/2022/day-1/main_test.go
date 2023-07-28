package main

import (
	"testing"

	"github.com/cassis163/advent-of-code/util"
)

func TestGetMostCaloriesCarriedByElf(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	want := 24000
	got := getMostCaloriesCarriedByElf(getSortedCaloriesPerElf(data))

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetCaloriesCarriedByTopThreeElves(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	want := 45000
	got := getCaloriesCarriedByTopThreeElves(getSortedCaloriesPerElf(data))

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
