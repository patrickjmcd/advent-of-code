package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() []string {
	var inputs []string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func part1() {
	inputs := readFile()

	var validPasswords []string
	// fmt.Printf("%+v", inputs)
	for _, passEntry := range inputs {
		r := regexp.MustCompile(`(?P<Min>\d*)\-(?P<Max>\d*)\s(?P<Char>.)\:\s(?P<Password>\w*)`)

		subMatch := r.FindStringSubmatch(passEntry)
		letter := subMatch[3]
		password := subMatch[4]
		occurences := strings.Count(password, letter)
		min, _ := strconv.Atoi(subMatch[1])
		max, _ := strconv.Atoi(subMatch[2])

		// fmt.Printf("min: %d, max: %d, occurrences: %d, letter: %s, password: %s\n", min, max, occurences, letter, password)

		if occurences >= min && occurences <= max {
			// fmt.Printf("min: %d, max: %d, occurrences: %d, letter: %s, password: %s\n", min, max, occurences, letter, password)
			validPasswords = append(validPasswords, password)
		}

	}

	fmt.Printf("# valid passwords = %d\n", len(validPasswords))
	fmt.Printf("# total passwords = %d\n", len(inputs))
}

func part2() {
	inputs := readFile()

	var validPasswords []string
	// fmt.Printf("%+v", inputs)
	for _, passEntry := range inputs {
		r := regexp.MustCompile(`(?P<Min>\d*)\-(?P<Max>\d*)\s(?P<Char>.)\:\s(?P<Password>\w*)`)

		subMatch := r.FindStringSubmatch(passEntry)
		letter := subMatch[3]
		password := subMatch[4]
		min, _ := strconv.Atoi(subMatch[1])
		max, _ := strconv.Atoi(subMatch[2])

		// fmt.Printf("min: %d, max: %d, occurrences: %d, letter: %s, password: %s\n", min, max, occurences, letter, password)

		validPositions := 0
		if string(password[min-1]) == letter {
			validPositions++
		}

		if string(password[max-1]) == letter {
			validPositions++
		}

		if validPositions == 1 {
			// fmt.Printf("min: %d, max: %d, letter: %s, password: %s\n", min, max, letter, password)
			validPasswords = append(validPasswords, password)
		}

	}

	fmt.Printf("# valid passwords = %d\n", len(validPasswords))
	fmt.Printf("# total passwords = %d\n", len(inputs))
}

func main() {
	part1()
	fmt.Println("\n-----------\nPART 2\n-----------")
	part2()
}
