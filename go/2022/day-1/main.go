package main

import (
	"fmt"
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

	sortedCaloriesPerElf := getSortedCaloriesPerElf(data)
	mostCaloriesCarriedByElf := getMostCaloriesCarriedByElf(sortedCaloriesPerElf)
	totalCaloriesCarriedByTopThreeElves := getCaloriesCarriedByTopThreeElves(sortedCaloriesPerElf)
	fmt.Printf("Most calories carried %d\n", mostCaloriesCarriedByElf)
	fmt.Printf("Total calories carried by top three %d\n", totalCaloriesCarriedByTopThreeElves)
}

func getSortedCaloriesPerElf(data string) []int {
	caloriesPerElf := make([]int, 0)

	currentIndex := 0
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			currentIndex++
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if len(caloriesPerElf) < currentIndex+1 {
			caloriesPerElf = append(caloriesPerElf, calories)
		} else {
			caloriesPerElf[currentIndex] += calories
		}
	}

	sort.Slice(caloriesPerElf, func(i, j int) bool { return caloriesPerElf[i] > caloriesPerElf[j] })

	return caloriesPerElf
}

func getMostCaloriesCarriedByElf(caloriesPerElf []int) int {
	return caloriesPerElf[0]
}

func getCaloriesCarriedByTopThreeElves(caloriesPerElf []int) int {
	return caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2]
}
