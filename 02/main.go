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

func getReports(input string) [][]int64 {
	var reports [][]int64
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\s+`)
	for _, line := range lines {
		numbers := strings.Split(re.ReplaceAllString(line, " "), " ")
		var levels []int64
		for i := range numbers {
			number, err := strconv.ParseInt(numbers[i], 10, 32)
			if err != nil {
				panic(err)
			}
			levels = append(levels, number)
		}
		reports = append(reports, levels)
	}
	return reports
}

func getAnswers(input string) (int64, int64) {
	reports := getReports(input)

	safeReports := int64(0)
	damperSafe := int64(0)
	for _, report := range reports {
		if checkSafety(report) {
			safeReports++
			damperSafe++
		} else {
			for i := 0; i < len(report); i++ {
				re := append(append([]int64{}, report[:i]...), report[i+1:]...)
				if checkSafety(re) {
					damperSafe++
					break
				}
			}
		}
	}

	return safeReports, damperSafe
}

func checkSafety(report []int64) bool {
	asc := report[0] < report[1]
	safe := true
	for i := 0; i < len(report)-1; i++ {
		if report[i] == report[i+1] {
			safe = false
		}
		if asc {
			if report[i] > report[i+1] {
				safe = false
			}
			if report[i+1]-report[i] > 3 {
				safe = false
			}
		} else {
			if report[i+1] > report[i] {
				safe = false
			}
			if report[i]-report[i+1] > 3 {
				safe = false
			}
		}
	}
	return safe
}
