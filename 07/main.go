package main

import (
	_ "embed"
	"fmt"
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

func getEquations(input string) map[int64][]int64 {
	var equations = map[int64][]int64{}
	gridRows := strings.Split(input, "\n")
	for _, row := range gridRows {
		split := strings.Split(row, ": ")
		parsedResult, _ := strconv.ParseInt(split[0], 10, 64)
		numbers := []int64{}
		for _, number := range strings.Split(split[1], " ") {
			parsedNumber, _ := strconv.ParseInt(number, 10, 64)
			numbers = append(numbers, parsedNumber)
		}
		equations[parsedResult] = numbers
	}
	return equations
}

func getAnswers(input string) (int64, int64) {
	equations := getEquations(input)

	totalCalibrationResult, totalCalibrationWithConcatenationResult := getTotalCalibrationResult(equations)

	return totalCalibrationResult, totalCalibrationWithConcatenationResult
}

func getTotalCalibrationResult(equations map[int64][]int64) (int64, int64) {
	totalCalibrationResult := int64(0)
	totalCalibrationWithConcatenationResult := int64(0)
	for result, numbers := range equations {
		if isEquationTrue(numbers, result) {
			totalCalibrationResult += result
		} else if isEquationWithConcatenationTrue(numbers, result) {
			totalCalibrationWithConcatenationResult += result
		}
	}
	return totalCalibrationResult, totalCalibrationResult + totalCalibrationWithConcatenationResult
}

func isEquationTrue(numbers []int64, result int64) bool {
	for i := 0; i < len(numbers)-1; i++ {
		multiplication := numbers[i] * numbers[i+1]
		sum := numbers[i] + numbers[i+1]
		if multiplication == result || sum == result {
			return true
		}

		if multiplication < result || sum < result {
			newMulNumbers := append([]int64{multiplication}, numbers[i+2:]...)
			newSumNumbers := append([]int64{sum}, numbers[i+2:]...)
			if isEquationTrue(newMulNumbers, result) || isEquationTrue(newSumNumbers, result) {
				return true
			}
		}
		return false
	}
	return false
}

func isEquationWithConcatenationTrue(numbers []int64, result int64) bool {
	for i := 0; i < len(numbers)-1; i++ {
		multiplication := numbers[i] * numbers[i+1]
		sum := numbers[i] + numbers[i+1]
		concatenatedNumber := fmt.Sprintf("%d%d", numbers[i], numbers[i+1])
		concatenation, _ := strconv.ParseInt(concatenatedNumber, 10, 64)
		if multiplication == result || sum == result || concatenation == result {
			return true
		}

		if multiplication < result || sum < result || concatenation < result {
			newMulNumbers := append([]int64{multiplication}, numbers[i+2:]...)
			newSumNumbers := append([]int64{sum}, numbers[i+2:]...)
			newConcatenationNumbers := append([]int64{concatenation}, numbers[i+2:]...)
			if isEquationWithConcatenationTrue(newMulNumbers, result) || isEquationWithConcatenationTrue(newSumNumbers, result) || isEquationWithConcatenationTrue(newConcatenationNumbers, result) {
				return true
			}
		}
		return false
	}
	return false
}
