package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
 * Sums adjacent integers in a file seperated by a newline
 * and reports the highest sum.
 */
func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mostCalories := 0
	cur := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if cur > mostCalories {
				mostCalories = cur
			}
			cur = 0
		} else {
			n, err := strconv.Atoi(line)
			check(err)
			cur += n
		}
	}

	// We have to check one last time in case the file doesn't end
	// with a newline and the last elf is carrying the most calories.
	if cur >= mostCalories {
		mostCalories = cur
	}

	// Check for a scanner error
	check(err)

	fmt.Printf("Answer: %d calories\n", mostCalories)
}
