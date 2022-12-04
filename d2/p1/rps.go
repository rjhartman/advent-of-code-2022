package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Score int
type MyStrategy rune
type OpponentStrategy rune

const (
	Win  Score = 6
	Loss       = 0
	Draw       = 3
)

// We can obtain the score for RPS values just by doing ascii math.
const OPPONENT_ROCK rune = 'A'
const ROCK rune = 'X'

// Play a round of rock paper scissors
func play(opponentPlay, myPlay rune) Score {
	o := opponentPlay - OPPONENT_ROCK
	m := myPlay - ROCK
	switch m - o {
	case 1, -2:
		return Win
	case -1, 2:
		return Loss
	default:
		return Draw
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		strategy := strings.Split(scanner.Text(), " ")
		opponentStrategy := []rune(strategy[0])[0]
		myStrategy := []rune(strategy[1])[0]
		score += int(myStrategy-ROCK) + int(play(opponentStrategy, myStrategy)) + 1
	}

	// Check for a scanner error
	check(err)

	fmt.Printf("Score: %d\n", score)
}
