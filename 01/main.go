package main

import (
	_ "embed"
	"fmt"
	"regexp"
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

func getLists(input string) ([]int64, []int64) {
	var leftList []int64
	var rightList []int64
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\s+`)
	for _, line := range lines {
		numbers := strings.Split(re.ReplaceAllString(line, " "), " ")
		for i := range numbers {
			number, err := strconv.ParseInt(numbers[i], 10, 32)
			if err != nil {
				panic(err)
			}
			if i == 0 {
				leftList = append(leftList, number)
			} else if i == 1 {
				rightList = append(rightList, number)
			} else {
				panic(fmt.Sprintf("More than two numbers in line %s", i))
			}
		}
	}
	return leftList, rightList
}

func getAnswers(input string) (int64, int64) {
	leftList, rightList := getLists(input)
	sort.Slice(leftList, func(i, j int) bool { return leftList[i] < leftList[j] })
	sort.Slice(rightList, func(i, j int) bool { return rightList[i] < rightList[j] })

	totalDistance := int64(0)
	similarity := int64(0)
	for i := range leftList {
		// First part
		if leftList[i] > rightList[i] {
			totalDistance += leftList[i] - rightList[i]
		} else {
			totalDistance += rightList[i] - leftList[i]
		}
		// Second part
		currentLeftNumber := leftList[i]
		same := 0
		for _, rightNumber := range rightList {
			if rightNumber == currentLeftNumber {
				same++
			}
		}
		similarity += currentLeftNumber * int64(same)
	}
	return totalDistance, similarity
}
