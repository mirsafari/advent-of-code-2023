package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += CardPoints(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

// Part one

func CardPoints(line string) int {
	score := 0.5

	// Remove prefix
	line = line[strings.IndexByte(line, ':')+2:]

	// Extract numbers into an array
	sets := strings.Split(line, "|")

	winningNumbers := []int{}
	myNumbers := []int{}

	// Covert to integers
	for _, number := range strings.Fields(sets[0]) {
		num, _ := strconv.Atoi(string(number))

		winningNumbers = append(winningNumbers, num)
	}

	for _, number := range strings.Fields(sets[1]) {
		num, _ := strconv.Atoi(string(number))

		myNumbers = append(myNumbers, num)
	}

	for _, x := range myNumbers {
		for _, y := range winningNumbers {
			if x == y {
				score *= 2
			}
		}
	}

	if score == 0.5 {
		return 0
	}
	return int(score)
}
