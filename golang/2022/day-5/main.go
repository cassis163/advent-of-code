package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/cassis163/advent-of-code/util"
)

func main() {
	data, err := util.ReadFileAsString("./data.txt")
	if err != nil {
		panic(err)
	}

	PartOne(data)
	PartTwo(data)
}

func PartOne(data string) string {
	supply := getSupplyFromData(data)

	moves := getMovesFromData(data)
	for _, move := range *moves {
		supply.ApplyCrateMover9000Move(move)
	}

	highestCrates := getHighestCrates(supply)
	fmt.Printf("Part one: %s\n", highestCrates)

	return highestCrates
}

func PartTwo(data string) string {
	supply := getSupplyFromData(data)

	moves := getMovesFromData(data)
	for _, move := range *moves {
		supply.ApplyCrateMover9001Move(move)
	}

	highestCrates := getHighestCrates(supply)
	fmt.Printf("Part two: %s\n", highestCrates)

	return highestCrates
}

type Move struct {
	fromStackIndex int
	toStackIndex   int
	amount         int
}

type Supply struct {
	stacks map[int]Stack
}

type Stack struct {
	crates []string
}

func (s *Supply) MoveWithCrateMover9000(fromStackIndex int, toStackIndex int, amount int) {
	// Move one crate at a time

	fromStack := s.stacks[fromStackIndex]
	toStack := s.stacks[toStackIndex]
	for i := 0; i < amount; i++ {
		fromCrateIndex := len(fromStack.crates) - 1
		fromCrate := fromStack.crates[fromCrateIndex]
		toStack.crates = append(toStack.crates, fromCrate)
		fromStack.crates = fromStack.crates[:fromCrateIndex]
	}
	s.stacks[fromStackIndex] = fromStack
	s.stacks[toStackIndex] = toStack
}

func (s *Supply) ApplyCrateMover9000Move(move Move) {
	s.MoveWithCrateMover9000(move.fromStackIndex-1, move.toStackIndex-1, move.amount)
}

func (s *Supply) MoveWithCrateMover9001(fromStackIndex int, toStackIndex int, amount int) {
	// Move all crates at once

	fromStack := s.stacks[fromStackIndex]
	toStack := s.stacks[toStackIndex]

	crates := fromStack.crates[len(fromStack.crates)-amount:]
	toStack.crates = append(toStack.crates, crates...)
	fromStack.crates = fromStack.crates[:len(fromStack.crates)-amount]

	s.stacks[fromStackIndex] = fromStack
	s.stacks[toStackIndex] = toStack
}

func (s *Supply) ApplyCrateMover9001Move(move Move) {
	s.MoveWithCrateMover9001(move.fromStackIndex-1, move.toStackIndex-1, move.amount)
}

func (s *Supply) AddCrateToStack(crate string, stackIndex int) {
	stack, ok := s.stacks[stackIndex]
	if !ok {
		stack = Stack{}
		s.stacks[stackIndex] = stack
	}

	stack.crates = append(stack.crates, crate)
	s.stacks[stackIndex] = stack
}

func getHighestCrates(supply *Supply) string {
	keys := make([]int, 0, len(supply.stacks))
	for k := range supply.stacks {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	highestCrates := ""
	for _, k := range keys {
		stack := supply.stacks[k]
		highestCrates += stack.crates[len(stack.crates)-1]
		fmt.Printf("%d: %s\n", k, stack.crates[len(stack.crates)-1])
	}

	return highestCrates
}

func getMovesFromData(data string) *[]Move {
	moves := []Move{}

	movesLines := strings.Split(data, "\n\n")[1]
	lines := strings.Split(movesLines, "\n")
	re := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		if line == "" {
			continue
		}

		digits := re.FindAllString(line, -1)
		fromStackIndex, err1 := strconv.Atoi(digits[1])
		toStackIndex, err2 := strconv.Atoi(digits[2])
		amount, err3 := strconv.Atoi(digits[0])

		if err := errors.Join(err1, err2, err3); err != nil {
			panic(err)
		}

		moves = append(moves, Move{
			fromStackIndex, toStackIndex, amount,
		})
	}

	return &moves
}

func getSupplyFromData(data string) *Supply {
	supply := &Supply{
		stacks: map[int]Stack{},
	}

	supplyLines := strings.Split(data, "\n\n")[0]
	lines := strings.Split(supplyLines, "\n")
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		crates := getCratesFromLine(line)
		for stackIndex, crate := range crates {
			supply.AddCrateToStack(crate, stackIndex)
		}
	}

	return supply
}

func getCratesFromLine(line string) map[int]string {
	crates := map[int]string{}
	stackIndex := 0
	for i := 1; i < len(line)-1; i += 4 {
		crate := string(line[i])
		if crate == " " {
			stackIndex++
			continue
		}

		crates[stackIndex] = crate
		stackIndex++
	}
	return crates
}
