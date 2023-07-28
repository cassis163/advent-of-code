package main

import (
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

const RockShapePlayer1 = "A"
const PaperShapePlayer1 = "B"
const ScissorsShapePlayer1 = "C"

const RockShapePlayer2 = "X"
const PaperShapePlayer2 = "Y"
const ScissorsShapePlayer2 = "Z"

const RockShape = "R"
const PaperShape = "P"
const ScissorsShape = "S"

const UnknownShapeError = "unknown shape"

type Turn struct {
	Player1Move string
	Player2Move string
}

func main() {
	data, err := util.ReadFileAsString("./data.txt")
	if err != nil {
		panic(err)
	}

	turns := getTurns(data)
	score := getScore(turns)
	fmt.Printf("Score: %d\n", score)
}

func getTurns(data string) []Turn {
	turns := make([]Turn, 0)

	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}

		moves := strings.Split(line, " ")
		turn := Turn{
			Player1Move: moves[0],
			Player2Move: moves[1],
		}
		turns = append(turns, turn)
	}

	return turns
}

func getScore(turns []Turn) int {
	totalScore := 0
	for _, turn := range turns {
		shapeScore := getShapeScore(turn.Player2Move)
		roundScore := getRoundScore(turn)
		totalScore += shapeScore + roundScore
	}

	return totalScore
}

func getRoundScore(turn Turn) int {
	shapes := []string{RockShape, PaperShape, ScissorsShape}
	player1Index, err := util.GetIndexInSlice(shapes, castPlayer1ShapeToShape(turn.Player1Move))
	if err != nil {
		panic(err)
	}

	player2Index, err := util.GetIndexInSlice(shapes, castPlayer2ShapeToShape(turn.Player2Move))
	if err != nil {
		panic(err)
	}

	if player1Index == player2Index {
		return DrawScore
	}
	if player1Index == player2Index-1 {
		return LossScore
	}
	return WinScore
}

func getShapeScore(shape string) int {
	switch shape {
	case RockShapePlayer2:
		return RockScore
	case PaperShapePlayer2:
		return PaperScore
	case ScissorsShapePlayer2:
		return ScissorsScore
	default:
		panic(UnknownShapeError)
	}
}

func castPlayer1ShapeToShape(player1Shape string) string {
	switch player1Shape {
	case RockShapePlayer1:
		return RockShape
	case PaperShapePlayer1:
		return PaperShape
	case ScissorsShapePlayer1:
		return ScissorsShape
	default:
		panic(UnknownShapeError)
	}
}

func castPlayer2ShapeToShape(player2Shape string) string {
	switch player2Shape {
	case RockShapePlayer2:
		return RockShape
	case PaperShapePlayer2:
		return PaperShape
	case ScissorsShapePlayer2:
		return ScissorsShape
	default:
		panic(UnknownShapeError)
	}
}
