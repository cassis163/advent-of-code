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

	assignmentPairs := getAssignmentPairs(data)

	want := 2
	got := getOverlappingAssignmentsCount(assignmentPairs)

	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}
