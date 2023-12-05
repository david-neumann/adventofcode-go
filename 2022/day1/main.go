package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func DayOne() {
	// Open the input file
	file, err := os.Open("2022/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Creates a scanner that will read the file line by line
	scanner := bufio.NewScanner(file)
	// A new slice to hold the count values
	var intSlice []int
	highCount, currentCount := 0, 0

	// Read the each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, that's the end of that elf's count
		// Compare count to highCount and set highCount, which is our answer for part 1
		// Add the count to our intSlice so we can track each elf's count
		if line == "" {
			if currentCount > highCount {
				highCount = currentCount
			}

			intSlice = append(intSlice, currentCount)
			currentCount = 0
			continue
		}

		// Convert string to int and add line to count
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		currentCount += num
	}

	// Sort the intSlice so we can get the 3 highest totals, which is our answer for part 2
	sort.Ints(intSlice)
	sumTopThree := intSlice[len(intSlice)-3] + intSlice[len(intSlice)-2] + intSlice[len(intSlice)-1]

	fmt.Println("Highest Count:", highCount)
	fmt.Println("The sum of the 3 largest totals is:", sumTopThree)
}
