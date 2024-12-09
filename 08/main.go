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
	antennas := getAntennas(grid)
	antinodes := getAntinodes(grid, antennas)
	antinodesWholeLine := getAntinodesWholeLine(grid, antennas)

	return len(antinodes), len(antinodesWholeLine)
}

func getAntinodes(grid [][]string, antennas map[string][][]int) map[string][]string {
	maxI := len(grid)
	maxJ := len(grid[0])
	antinodes := map[string][]string{}
	for antenna, locations := range antennas {
		for i := range locations {
			for j := range locations {
				if i == j {
					continue
				}
				firstAntenna := locations[i]
				secondAntenna := locations[j]

				firstI := 2*firstAntenna[0] - secondAntenna[0]
				firstJ := 2*firstAntenna[1] - secondAntenna[1]
				if firstI >= 0 && firstI < maxI && firstJ >= 0 && firstJ < maxJ {
					firstAntinode := fmt.Sprintf("%d,%d", firstI, firstJ)
					if antinodes[firstAntinode] == nil {
						antinodes[firstAntinode] = []string{}
					}
					antinodes[firstAntinode] = append(antinodes[firstAntinode], antenna)
				}

				secondI := 2*secondAntenna[0] - firstAntenna[0]
				secondJ := 2*secondAntenna[1] - firstAntenna[1]
				if secondI >= 0 && secondI < maxI && secondJ >= 0 && secondJ < maxJ {
					secondAntinode := fmt.Sprintf("%d,%d", secondI, secondJ)
					if antinodes[secondAntinode] == nil {
						antinodes[secondAntinode] = []string{}
					}
					antinodes[secondAntinode] = append(antinodes[secondAntinode], antenna)
				}
			}
		}
	}
	return antinodes
}

func getAntinodesWholeLine(grid [][]string, antennas map[string][][]int) map[string][]string {
	maxI := len(grid)
	maxJ := len(grid[0])
	antinodes := map[string][]string{}
	for antenna, locations := range antennas {
		for i := range locations {
			for j := range locations {
				if i == j {
					continue
				}
				firstAntenna := locations[i]
				secondAntenna := locations[j]

				for mul := 0; true; mul++ {
					firstI := (mul+1)*firstAntenna[0] - mul*secondAntenna[0]
					firstJ := (mul+1)*firstAntenna[1] - mul*secondAntenna[1]
					if firstI >= 0 && firstI < maxI && firstJ >= 0 && firstJ < maxJ {
						firstAntinode := fmt.Sprintf("%d,%d", firstI, firstJ)
						if antinodes[firstAntinode] == nil {
							antinodes[firstAntinode] = []string{}
						}
						antinodes[firstAntinode] = append(antinodes[firstAntinode], antenna)
					} else {
						break
					}
				}

				for mul := 0; true; mul++ {
					secondI := (mul+1)*secondAntenna[0] - mul*firstAntenna[0]
					secondJ := (mul+1)*secondAntenna[1] - mul*firstAntenna[1]
					if secondI >= 0 && secondI < maxI && secondJ >= 0 && secondJ < maxJ {
						secondAntinode := fmt.Sprintf("%d,%d", secondI, secondJ)
						if antinodes[secondAntinode] == nil {
							antinodes[secondAntinode] = []string{}
						}
						antinodes[secondAntinode] = append(antinodes[secondAntinode], antenna)
					} else {
						break
					}
				}
			}
		}
	}
	return antinodes
}

func getAntennas(grid [][]string) map[string][][]int {
	antennas := map[string][][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != "." {
				if antennas[grid[i][j]] == nil {
					antennas[grid[i][j]] = [][]int{}
				}
				antennas[grid[i][j]] = append(antennas[grid[i][j]], []int{i, j})
			}
		}
	}
	return antennas
}
