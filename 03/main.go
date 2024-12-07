package main

import (
	_ "embed"
	"fmt"
	"regexp"
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

func getAnswers(input string) (int64, int64) {
	mulPattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`
	mulRegex := regexp.MustCompile(mulPattern)
	mulMatchIndexes := mulRegex.FindAllStringIndex(input, -1)

	doPattern := `do\(\)`
	doRegex := regexp.MustCompile(doPattern)
	doMatchIndexes := doRegex.FindAllStringIndex(input, -1)

	doNotPattern := `don't\(\)`
	doNotRegex := regexp.MustCompile(doNotPattern)
	doNotMatchIndexes := doNotRegex.FindAllStringIndex(input, -1)

	multiplications := int64(0)
	enabledMultiplications := int64(0)
	for _, match := range mulMatchIndexes {
		multipliers := strings.Replace(input[match[0]:match[1]], "mul(", "", 1)
		multipliers = strings.Replace(multipliers, ")", "", 1)
		split := strings.Split(multipliers, ",")
		first, _ := strconv.ParseInt(split[0], 10, 64)
		second, _ := strconv.ParseInt(split[1], 10, 64)
		multiplications += first * second

		lastDo := -1
		lastDoNot := -1
		for _, j := range doMatchIndexes {
			if match[0] > j[0] {
				lastDo = j[0]
			} else {
				break
			}
		}
		for _, j := range doNotMatchIndexes {
			if match[0] > j[0] {
				lastDoNot = j[0]
			} else {
				break
			}
		}
		if lastDoNot == -1 || lastDo > lastDoNot {
			enabledMultiplications += first * second
		}
	}

	return multiplications, enabledMultiplications
}
