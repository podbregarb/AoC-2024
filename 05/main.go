package main

import (
	"aoc-2024/utils"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	fmt.Println("Test input")
	printAnswers(testInput)
	fmt.Println("Input")
	printAnswers(input)
}

func printAnswers(s string) {
	first, second := getAnswers(s)
	fmt.Printf("First answer: %d\n", first)
	fmt.Printf("Second answer: %d\n", second)
}

func getRules(input string) map[int64][]int64 {
	rules := make(map[int64][]int64)
	rulesRows := strings.Split(strings.Split(input, "\n\n")[0], "\n")
	for _, row := range rulesRows {
		numbers := strings.Split(row, "|")
		firstPage, _ := strconv.ParseInt(numbers[0], 10, 64)
		secondPage, _ := strconv.ParseInt(numbers[1], 10, 64)
		if rules[firstPage] == nil {
			rules[firstPage] = []int64{}
		}
		rules[firstPage] = append(rules[firstPage], secondPage)
	}
	return rules
}

func getUpdates(input string) [][]int64 {
	var updates [][]int64
	updatesRows := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	for _, row := range updatesRows {
		var update []int64
		numbers := strings.Split(row, ",")
		for _, number := range numbers {
			parsed, _ := strconv.ParseInt(number, 10, 64)
			update = append(update, parsed)
		}
		updates = append(updates, update)
	}
	return updates
}

func getAnswers(input string) (int64, int64) {
	rules := getRules(input)
	updates := getUpdates(input)

	orderedUpdatesSum := int64(0)
	unorderedUpdatesSum := int64(0)
	for _, update := range updates {
		if isOrdered(rules, update) {
			orderedUpdatesSum += update[len(update)/2]
		} else {
			orderedUpdate := orderUpdate(rules, update)
			unorderedUpdatesSum += orderedUpdate[len(update)/2]
		}
	}
	return orderedUpdatesSum, unorderedUpdatesSum
}

func isOrdered(rules map[int64][]int64, update []int64) bool {
	for i := range update {
		for j := i + 1; j < len(update); j++ {
			if !utils.IntContains(rules[update[i]], update[j]) {
				return false
			}
		}
	}
	return true
}

func orderUpdate(rules map[int64][]int64, update []int64) []int64 {
	sort.Slice(update, func(i, j int) bool {
		if utils.IntContains(rules[update[i]], update[j]) {
			return true
		}
		return false
	})
	return update
}
