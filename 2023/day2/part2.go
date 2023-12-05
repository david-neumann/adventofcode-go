package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartTwo() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sum := 0

	for sc.Scan() {
		setPower := findSetPower(sc.Text())
		sum += setPower
	}

	fmt.Println("Sum:", sum)
}

func findSetPower(line string) int {
	minCubesNeeded := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	regExp := regexp.MustCompile(`\d+\s+\w+`)
	matches := regExp.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		split := strings.Split(match[0], " ")
		count, _ := strconv.Atoi(split[0])
		color := split[1]
		if count > minCubesNeeded[color] {
			minCubesNeeded[color] = count
		}
	}

	setPower := minCubesNeeded["red"] * minCubesNeeded["green"] * minCubesNeeded["blue"]

	return setPower
}
