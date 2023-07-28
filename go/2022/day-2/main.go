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

func main() {
	data, err := util.ReadFileAsString("./data.txt")
	if err != nil {
		panic(err)
	}

	turns := getTurns(data)
	score := getScore(turns)
	fmt.Printf("Score: %d\n", score)
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

func getScore(turns []string) int {
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
