package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFilePath = "2022/day3/input.txt"

func PartOne() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sum := 0

	for sc.Scan() {
		rucksack := sc.Text()
		midpoint := len(rucksack) / 2
		compartment1, compartment2 := rucksack[:midpoint], rucksack[midpoint:]
		commonItem := findCommon(compartment1, compartment2)
		itemValue := getItemValue(commonItem)
		sum += itemValue
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func findCommon(one, two string) rune {
	for _, letter := range one {
		if strings.ContainsRune(two, letter) {
			return letter
		}
	}

	return 0
}

func getItemValue(item rune) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterValues := make(map[rune]int)

	for i, letter := range alphabet {
		letterValues[letter] = i + 1
	}

	itemValue := letterValues[item]
	return itemValue
}

// 1. Go through each line of the input
// 2. Get the length of the line
// 3. Split the line in half
// 4. Find the character that appears in both halves
// 5. Get the numeric value of the shared character
// 6. Add that value to a running total
// 7. Return the total
