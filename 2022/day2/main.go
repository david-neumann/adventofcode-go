package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFilePath = "2022/day2/input.txt"

type gameElem int

const (
	rock gameElem = iota
	paper
	scissors
)

var charToElem = map[byte]gameElem{
	'A': rock,
	'B': paper,
	'C': scissors,
	'X': rock,
	'Y': paper,
	'Z': scissors,
}

var shapeScore = map[gameElem]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

const (
	winPoints  = 6
	drawPoints = 3
	losePoints = 0
)

var winningPlays = map[gameElem]gameElem{
	rock:     paper,
	paper:    scissors,
	scissors: rock,
}

var losingPlays = map[gameElem]gameElem{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

type outcomeType int

const (
	winType outcomeType = iota
	drawType
	loseType
)

var charToOutcome = map[byte]outcomeType{
	'X': loseType,
	'Y': drawType,
	'Z': winType,
}

func DayTwo() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	totalScore := processGame(sc)

	fmt.Println(totalScore)
}

func processGame(sc *bufio.Scanner) int {
	totalScore := 0

	for sc.Scan() {
		round := sc.Text()
		opp, youShould := round[0], round[2]
		roundScore := determineRoundOutcome(opp, youShould)
		totalScore += roundScore
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return totalScore
}

func determineRoundOutcome(opp, youShould byte) int {
	oppShape := charToElem[opp]
	outcome := charToOutcome[youShould]
	var yourShape gameElem
	roundScore := 0

	switch outcome {
	case winType:
		yourShape = winningPlays[oppShape]
		roundScore = winPoints + shapeScore[yourShape]
	case loseType:
		yourShape = losingPlays[oppShape]
		roundScore = losePoints + shapeScore[yourShape]
	case drawType:
		yourShape = oppShape
		roundScore = drawPoints + shapeScore[yourShape]
	}

	return roundScore
}
