package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cassis163/advent-of-code/util"
	"golang.org/x/exp/slices"
)

type AssignmentPair struct {
	assignmentsA, assignmentsB []int
}

func main() {
	data, err := util.ReadFileAsString("./data.txt")
	if err != nil {
		panic(err)
	}

	assignmentPairs := getAssignmentPairs(data)
	count := getOverlappingAssignmentsCount(assignmentPairs)

	fmt.Printf("There are %d overlapping assignments\n", count)
}

func getOverlappingAssignmentsCount(assignmentPairs []AssignmentPair) int {
	count := 0
	for _, assignmentPair := range assignmentPairs {
		if assignmentPairOverlaps(assignmentPair) {
			count++
		}
	}

	return count
}

func assignmentPairOverlaps(assignmentPair AssignmentPair) bool {
	pairA := assignmentPair.assignmentsA
	pairB := assignmentPair.assignmentsB

	return doesIntArrayContain(pairA, pairB) || doesIntArrayContain(pairB, pairA)
}

func doesIntArrayContain(container []int, contained []int) bool {
	for _, containedInt := range contained {
		if !slices.Contains(container, containedInt) {
			return false
		}
	}

	return true
}

func getAssignmentPairs(data string) []AssignmentPair {
	var assignmentPairs []AssignmentPair
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}

		assignmentPairs = append(assignmentPairs, getAssignmentPair(line))
	}

	return assignmentPairs
}

func getAssignmentPair(line string) AssignmentPair {
	assignments := strings.Split(line, ",")
	assignmentsA := getAssignmentNumbers(assignments[0])
	assignmentsB := getAssignmentNumbers(assignments[1])

	return AssignmentPair{
		assignmentsA: assignmentsA,
		assignmentsB: assignmentsB,
	}
}

func getAssignmentNumbers(assignments string) []int {
	assignmentsRepresentation := strings.Split(assignments, "-")
	fromAssignment, err := strconv.Atoi(assignmentsRepresentation[0])
	if err != nil {
		panic(err)
	}

	toAssignment, err := strconv.Atoi(assignmentsRepresentation[1])
	if err != nil {
		panic(err)
	}

	var assignmentNumbers []int
	for i := fromAssignment; i <= toAssignment; i++ {
		assignmentNumbers = append(assignmentNumbers, i)
	}

	return assignmentNumbers
}
