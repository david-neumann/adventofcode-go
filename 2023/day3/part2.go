package day3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func PartTwo(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	grid := convertToGrid(lines)
	gearRatios := findGears(grid)
	sumOfGearRatios := findSum(gearRatios)

	return sumOfGearRatios, nil
}

func findGears(grid [][]rune) []int {
	uniqueNumbers := make(map[int]map[position]struct{})
	symbolNumbers := make(map[position]map[int]struct{})

	directions := [][]int{
		{0, -1},  // Left
		{-1, -1}, // Up-Left
		{-1, 0},  // Up
		{-1, 1},  // Up-Right
		{0, 1},   // Right
		{1, 1},   // Down-Right
		{1, 0},   // Down
		{1, -1},  // Down-Left
	}

	// Iterate through each row in the grid
	for i, row := range grid {
		// Iterate through each character in the row
		for j, char := range row {
			if isSymbol(char) {
				// Check all adjacent cells when a symbol is found
				for _, dir := range directions {
					newRow, newCol := i+dir[0], j+dir[1]
					adjacentCell := grid[newRow][newCol]
					// Check if the adjactent cell is a number
					if unicode.IsDigit(adjacentCell) {
						// Find the full number in the row
						foundNum := findConsecutiveNumbers(grid[newRow], newCol)
						symbolPosition := position{row: i, col: j}

						if uniqueNumbers[foundNum] == nil {
							uniqueNumbers[foundNum] = make(map[position]struct{})
						}

						uniqueNumbers[foundNum][symbolPosition] = struct{}{}

						if symbolNumbers[symbolPosition] == nil {
							symbolNumbers[symbolPosition] = make(map[int]struct{})
						}

						symbolNumbers[symbolPosition][foundNum] = struct{}{}
					}
				}
			}
		}
	}

	var gearRatios []int
	for position, numbers := range symbolNumbers {
		if grid[position.row][position.col] == '*' && len(numbers) == 2 {
			var positionNums []int
			for num := range numbers {
				positionNums = append(positionNums, num)
			}
			gearRatio := positionNums[0] * positionNums[1]
			gearRatios = append(gearRatios, gearRatio)
		}
	}

	return gearRatios
}
