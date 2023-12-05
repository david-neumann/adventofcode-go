package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type position struct {
	row, col int
}

func PartOne(filepath string) (int, error) {
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
	adjacentNumbers := findAdjacentNumbers(grid)
	sumOfAdjacent := findSum(adjacentNumbers)

	return sumOfAdjacent, nil
}

func convertToGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func findAdjacentNumbers(grid [][]rune) []int {
	uniqueNumbers := make(map[int]map[position]struct{})

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
					}
				}
			}
		}
	}

	// Convert map to slice of the found numbers
	var adjacentNumbers []int
	for num, positions := range uniqueNumbers {
		for range positions {
			adjacentNumbers = append(adjacentNumbers, num)
		}
	}

	return adjacentNumbers
}

func findConsecutiveNumbers(row []rune, startCol int) int {
	// Iterate to the left
	leftIndex := startCol
	for leftIndex > 0 && unicode.IsDigit(row[leftIndex-1]) {
		leftIndex--
	}

	// Iterate to the right
	rightIndex := startCol
	for rightIndex < len(row)-1 && unicode.IsDigit(row[rightIndex+1]) {
		rightIndex++
	}

	// Convert to an int
	foundNum, _ := strconv.Atoi(string(row[leftIndex : rightIndex+1]))

	return foundNum
}

func findSum(adjacentNumbers []int) int {
	sum := 0

	for _, number := range adjacentNumbers {
		sum += number
	}

	return sum
}

func isSymbol(char rune) bool {
	return char != '.' && !unicode.IsSpace(char) && !unicode.IsDigit(char)
}

// 1. Iterate over the entire grid searching for non-'.' symbols
// 2. Once a symbol is found, check all adjacent cells looking for a number
// 3. If if a number is found, use a regex on that row to match the full number at that position
// 4. Add the found number to an []int
// 5. After the entire grid has been searched, sum up all of the numbers in the found []int
