package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cassis163/advent-of-code/util"
)

const RockScore = 1
const PaperScore = 2
const ScissorsScore = 3

const LossScore = 0
const DrawScore = 3
const WinScore = 6

var Wins = [3]string{"AY", "BZ", "CX"}
var Draws = [3]string{"AX", "BY", "CZ"}
var Losses = [3]string{"AZ", "BX", "CY"}

var InstructionsMap = map[string]string{
	"AX": "AZ",
	"AY": "AX",
	"AZ": "AY",
	"BX": "BX",
	"BY": "BY",
	"BZ": "BZ",
	"CX": "CY",
	"CY": "CZ",
	"CZ": "CX",
}

func main() {
	data, err := util.ReadFileAsString("./data.txt")
	if err != nil {
		panic(err)
	}

	turns := getTurns(data)
	score := getScoreForPartOne(turns)
	fmt.Printf("Score (part 1): %d\n", score)
	score2 := getScoreForPartTwo(turns)
	fmt.Printf("Score (part 2): %d\n", score2)
}

func getTurns(data string) []string {
	turns := make([]string, 0)
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}

		moves := strings.Split(line, " ")
		turns = append(turns, moves[0]+moves[1])
	}

	return turns
}

func getScoreForPartOne(turns []string) int {
	totalScore := 0
	for _, turn := range turns {
		shapeScore, err := getShapeScore(turn)
		if err != nil {
			panic(err)
		}

		roundScore, err := getRoundScore(turn)
		if err != nil {
			panic(err)
		}

		totalScore += shapeScore + roundScore
	}

	return totalScore
}

func getScoreForPartTwo(turns []string) int {
	totalScore := 0
	for _, turnWithInstruction := range turns {
		turn := upcastTurnWithInstruction(turnWithInstruction)
		shapeScore, err := getShapeScore(turn)
		if err != nil {
			panic(err)
		}

		roundScore, err := getRoundScore(turn)
		if err != nil {
			panic(err)
		}

		totalScore += shapeScore + roundScore
	}

	return totalScore
}

func getShapeScore(turn string) (int, error) {
	move := turn[1:2]
	switch move {
	case "X":
		return RockScore, nil
	case "Y":
		return PaperScore, nil
	case "Z":
		return ScissorsScore, nil
	default:
		return 0, errors.New("unknown shape")
	}
}

func getRoundScore(turn string) (int, error) {
	if util.IsValueInSlice(Wins[:], turn) {
		return WinScore, nil
	}

	if util.IsValueInSlice(Draws[:], turn) {
		return DrawScore, nil
	}

	if util.IsValueInSlice(Losses[:], turn) {
		return LossScore, nil
	}

	return 0, errors.New("missing round score")
}

func upcastTurnWithInstruction(turn string) string {
	return InstructionsMap[turn[0:2]]
}
