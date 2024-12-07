package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

var XMAS = "XMAS"
var MAS = "MAS"

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

func getChars(input string) [][]string {
	var chars [][]string
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		rowChars := strings.Split(row, "")
		chars = append(chars, rowChars)
	}
	return chars
}

func getAnswers(input string) (int, int) {
	chars := getChars(input)

	xmasCount := countXmas(chars)

	xMasCount := countxMas(chars)

	return xmasCount, xMasCount
}

func countXmas(chars [][]string) int {
	xmasCount := 0
	for i := range chars {
		for j := range chars[i] {
			if chars[i][j] == "X" {
				if isXmas(chars, i, j, 0, 1, chars[i][j], XMAS) { // Right
					xmasCount++
				}
				if isXmas(chars, i, j, 0, -1, chars[i][j], XMAS) { // Left
					xmasCount++
				}
				if isXmas(chars, i, j, 1, 0, chars[i][j], XMAS) { // Down
					xmasCount++
				}
				if isXmas(chars, i, j, -1, 0, chars[i][j], XMAS) { // Up
					xmasCount++
				}
				if isXmas(chars, i, j, -1, -1, chars[i][j], XMAS) { // Left Up
					xmasCount++
				}
				if isXmas(chars, i, j, 1, -1, chars[i][j], XMAS) { // Left Down
					xmasCount++
				}
				if isXmas(chars, i, j, -1, 1, chars[i][j], XMAS) { // Right Up
					xmasCount++
				}
				if isXmas(chars, i, j, 1, 1, chars[i][j], XMAS) { // Right Down
					xmasCount++
				}
			}
		}
	}
	return xmasCount
}

func isXmas(chars [][]string, i int, j int, iDir int, jDir int, word string, searchWord string) bool {
	iNew := i + iDir
	jNew := j + jDir
	if (iNew < 0 || iNew >= len(chars)) || (jNew < 0 || jNew >= len(chars[0])) {
		return false
	}
	word += chars[iNew][jNew]
	if word == searchWord {
		return true
	}
	if strings.HasPrefix(searchWord, word) {
		return isXmas(chars, iNew, jNew, iDir, jDir, word, searchWord)
	}
	return false
}

func countxMas(chars [][]string) int {
	xmasCount := 0
	for i := range chars {
		for j := range chars[i] {
			if chars[i][j] == "A" {
				if i-1 >= 0 && i+1 < len(chars) && j-1 >= 0 && j+1 < len(chars[i]) {
					isDownRightMas := isXmas(chars, i-1, j-1, 1, 1, chars[i-1][j-1], MAS)
					isUpLeftMas := isXmas(chars, i+1, j+1, -1, -1, chars[i+1][j+1], MAS)
					isUpRightMas := isXmas(chars, i+1, j-1, -1, 1, chars[i+1][j-1], MAS)
					isDownLeftMas := isXmas(chars, i-1, j+1, 1, -1, chars[i-1][j+1], MAS)
					if (isDownRightMas || isUpLeftMas) &&
						(isUpRightMas || isDownLeftMas) {
						xmasCount++
					}
				}
			}
		}
	}
	return xmasCount
}
