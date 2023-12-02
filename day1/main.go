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
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	sumWithLetters := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += SumOfSingleLine(scanner.Text())
		sumWithLetters += SumOfSingleLineWithLetters(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
	fmt.Println(sumWithLetters)
}

// Part one
func SumOfSingleLine(line string) int {

	var numbers []string

	for _, letter := range line {
		_, err := strconv.Atoi(string(letter))

		if err != nil {
			continue
		}

		numbers = append(numbers, string(letter))
	}

	newNumber, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])

	return newNumber
}

// Part two

var nameToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func SumOfSingleLineWithLetters(line string) int {

	leftDigit := 0
	rightDigit := 0

	for i, char := range line {
		num, err := strconv.Atoi(string(char))

		// It is a char
		if err != nil {

			for k, v := range nameToNumber {
				if strings.HasPrefix(line[i:], k) {
					if leftDigit == 0 {
						leftDigit = v
						rightDigit = v
					} else {
						rightDigit = v
					}
				}
			}

			continue
		}

		if leftDigit == 0 {
			leftDigit = num
			rightDigit = num
		} else {
			rightDigit = num
		}
	}

	newNumberString := strconv.Itoa(leftDigit) + strconv.Itoa(rightDigit)
	newNumber, _ := strconv.Atoi(newNumberString)
	return newNumber
}
