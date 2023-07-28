package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cassis163/advent-of-code/util"
)

func main() {
	data, err := util.ReadFileAsString("data.txt")
	if err != nil {
		panic(err)
	}

	mostCaloriesCarriedByElf := getMostCaloriesCarriedByElf(data)
	fmt.Println(mostCaloriesCarriedByElf)
}

func getMostCaloriesCarriedByElf(data string) int {
	lines := strings.Split(data, "\n")

	mostCalories := 0
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		currentCalories += calories
		if currentCalories > mostCalories {
			mostCalories = currentCalories
		}
	}

	return mostCalories
}
