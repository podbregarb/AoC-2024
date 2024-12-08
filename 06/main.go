package main

import (
	"aoc-2024/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

var DIRECTIONS = []string{">", "<", "^", "v"}

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

func getGrid(input string) [][]string {
	var grid [][]string
	gridRows := strings.Split(input, "\n")
	for _, row := range gridRows {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}

func getAnswers(input string) (int, int) {
	grid := getGrid(input)

	i, j := getGuardLocation(grid)
	gridCopy := utils.CopySlice(grid)
	walkedGrid := walk(gridCopy, i, j)
	uniqueLocations := countUniqueLocations(walkedGrid)

	obstacles := countObstacles(walkedGrid, grid, i, j)

	return uniqueLocations, obstacles
}

func countObstacles(walkedGrid [][]string, grid [][]string, iG int, jG int) int {
	obstacles := 0
	for i := range walkedGrid {
		for j := range walkedGrid[i] {
			if i == iG && j == jG {
				continue
			} else if walkedGrid[i][j] == "X" {
				gridCopy := utils.CopySlice(grid)
				gridCopy[i][j] = "#"
				if _, isCircular := isWalkCircular(gridCopy, iG, jG, map[string][]string{}); isCircular {
					obstacles++
				}
			}
		}
	}
	return obstacles
}

func countUniqueLocations(grid [][]string) int {
	uniqueLocations := 0
	for _, row := range grid {
		for _, char := range row {
			if char == "X" {
				uniqueLocations++
			}
		}
	}
	return uniqueLocations
}

func walk(grid [][]string, i int, j int) [][]string {
	direction := grid[i][j]
	if direction == ">" {
		grid[i][j] = "X"
		jNew := j + 1
		if jNew > len(grid[i])-1 {
			return grid
		}
		next := grid[i][jNew]
		if next == "." || next == "X" {
			grid[i][jNew] = ">"
			return walk(grid, i, jNew)
		} else {
			grid[i][j] = "v"
			return walk(grid, i, j)
		}
	} else if direction == "<" {
		grid[i][j] = "X"
		jNew := j - 1
		if jNew < 0 {
			return grid
		}
		next := grid[i][jNew]
		if next == "." || next == "X" {
			grid[i][jNew] = "<"
			return walk(grid, i, jNew)
		} else {
			grid[i][j] = "^"
			return walk(grid, i, j)
		}
	} else if direction == "^" {
		grid[i][j] = "X"
		iNew := i - 1
		if iNew < 0 {
			return grid
		}
		next := grid[iNew][j]
		if next == "." || next == "X" {
			grid[iNew][j] = "^"
			return walk(grid, iNew, j)
		} else {
			grid[i][j] = ">"
			return walk(grid, i, j)
		}
	} else if direction == "v" {
		grid[i][j] = "X"
		iNew := i + 1
		if iNew > len(grid)-1 {
			return grid
		}
		next := grid[iNew][j]
		if next == "." || next == "X" {
			grid[iNew][j] = "v"
			return walk(grid, iNew, j)
		} else {
			grid[i][j] = "<"
			return walk(grid, i, j)
		}
	} else {
		return grid
	}
}
func isWalkCircular(grid [][]string, i int, j int, directions map[string][]string) ([][]string, bool) {
	direction := grid[i][j]
	key := fmt.Sprintf("%d,%d", i, j)
	if directions[key] == nil {
		directions[key] = []string{}
	}
	if utils.StringContains(directions[key], direction) {
		return grid, true
	}
	directions[key] = append(directions[key], direction)

	if direction == ">" {
		grid[i][j] = "X"
		jNew := j + 1
		if jNew > len(grid[i])-1 {
			return grid, false
		}
		next := grid[i][jNew]
		if next == "." || next == "X" {
			grid[i][jNew] = ">"
			return isWalkCircular(grid, i, jNew, directions)
		} else {
			grid[i][j] = "v"
			return isWalkCircular(grid, i, j, directions)
		}
	} else if direction == "<" {
		grid[i][j] = "X"
		jNew := j - 1
		if jNew < 0 {
			return grid, false
		}
		next := grid[i][jNew]
		if next == "." || next == "X" {
			grid[i][jNew] = "<"
			return isWalkCircular(grid, i, jNew, directions)
		} else {
			grid[i][j] = "^"
			return isWalkCircular(grid, i, j, directions)
		}
	} else if direction == "^" {
		grid[i][j] = "X"
		iNew := i - 1
		if iNew < 0 {
			return grid, false
		}
		next := grid[iNew][j]
		if next == "." || next == "X" {
			grid[iNew][j] = "^"
			return isWalkCircular(grid, iNew, j, directions)
		} else {
			grid[i][j] = ">"
			return isWalkCircular(grid, i, j, directions)
		}
	} else if direction == "v" {
		grid[i][j] = "X"
		iNew := i + 1
		if iNew > len(grid)-1 {
			return grid, false
		}
		next := grid[iNew][j]
		if next == "." || next == "X" {
			grid[iNew][j] = "v"
			return isWalkCircular(grid, iNew, j, directions)
		} else {
			grid[i][j] = "<"
			return isWalkCircular(grid, i, j, directions)
		}
	} else {
		return grid, false
	}
}

func getGuardLocation(grid [][]string) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if utils.StringContains(DIRECTIONS, grid[i][j]) {
				return i, j
			}
		}
	}
	return -1, -1
}
