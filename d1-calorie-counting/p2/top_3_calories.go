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

// A reverse-sorted integer array that supports O(n) in-place inserts
type ElfList []int

func (el ElfList) Insert(n int) {
	for i, c := range el {
		if n > c {
			// Insert value and push all others down
			prev := n
			for j := i; j < len(el); j++ {
				prev, el[j] = el[j], prev
			}
			return
		}
	}
}

func (el ElfList) Sum() int {
	sum := 0
	for _, c := range el {
		sum += c
	}
	return sum
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list := ElfList{0, 0, 0}

	cur := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if cur > list[len(list)-1] {
				list.Insert(cur)
				fmt.Printf("Found %d. Have %v\n", cur, list)
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
	if cur >= list[len(list)-1] {
		list.Insert(cur)
	}

	// Check for a scanner error
	check(err)

	fmt.Printf("Top 3: %v\n", list)
	fmt.Printf("Answer: %d calories\n", list.Sum())
}
