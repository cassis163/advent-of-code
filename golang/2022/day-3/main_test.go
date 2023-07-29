package main

import (
	"testing"

	"github.com/cassis163/advent-of-code/util"
)

const ErrorMessage = "got %d want %d"

func TestPart1(t *testing.T) {
	data, err := util.ReadFileAsString("./test-data.txt")
	if err != nil {
		panic(err)
	}

	rucksacks := getRucksacks(data)
	foundItemTypes := getItemsInRucksacks(rucksacks)

	want := 157
	got := getPrioritySum(foundItemTypes)

	if got != want {
		t.Errorf(ErrorMessage, got, want)
	}
}

func TestGetPriorityLowercaseA(t *testing.T) {
	want := 1
	got := getPriorityOfItemType("a")

	if got != want {
		t.Errorf(ErrorMessage, got, want)
	}
}

func TestGetPriorityUpperCaseA(t *testing.T) {
	want := 27
	got := getPriorityOfItemType("A")

	if got != want {
		t.Errorf(ErrorMessage, got, want)
	}
}

func TestGetPriorityLowercaseZ(t *testing.T) {
	want := 26
	got := getPriorityOfItemType("z")

	if got != want {
		t.Errorf(ErrorMessage, got, want)
	}
}

func TestGetPriorityUpperCaseZ(t *testing.T) {
	want := 52
	got := getPriorityOfItemType("Z")

	if got != want {
		t.Errorf(ErrorMessage, got, want)
	}
}
