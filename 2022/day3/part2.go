package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo() {
	lines, err := readLines(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	groups := makeGroupsOfThree(lines)
	badges := findBadges(groups)
	badgeValues := calcBadgeValues(badges)
	totalValue := calcTotal(badgeValues)

	fmt.Println(totalValue)
}

func readLines(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func makeGroupsOfThree(lines []string) [][]string {
	var groups [][]string

	for i := 0; i < len(lines)-2; i += 3 {
		group := []string{lines[i], lines[i+1], lines[i+2]}
		groups = append(groups, group)
	}

	return groups
}

func findBadges(groups [][]string) string {
	var badges strings.Builder

outerLoop:
	for _, group := range groups {
		for _, item := range group[0] {
			if strings.ContainsRune(group[1], item) && strings.ContainsRune(group[2], item) {
				badges.WriteString(string(item))
				continue outerLoop
			}
		}
	}

	fmt.Println("badges count inside findBadges:", len(badges.String()))

	return badges.String()
}

func calcBadgeValues(badges string) []int {
	var badgeValues []int

	for _, item := range badges {
		itemValue := getItemValue(item)
		badgeValues = append(badgeValues, itemValue)
	}

	return badgeValues
}

func calcTotal(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}
