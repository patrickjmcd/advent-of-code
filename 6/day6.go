package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

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
	rows := readFile()
	input := strings.Join(rows, "|")
	answers := strings.Split(input, "||")
	countAllAnswers := 0
	for _, answer := range answers {
		answer1Line := strings.ReplaceAll(answer, "|", "")
		// fmt.Printf("%s - ", answer1Line)

		uniqueAnswers := ""
		for _, letter := range answer1Line {
			if !strings.Contains(uniqueAnswers, string(letter)) {
				uniqueAnswers += string(letter)
			}
		}
		// fmt.Printf("%s - ", uniqueAnswers)
		countAllAnswers += len(uniqueAnswers)
		// fmt.Printf("%d\n", len(uniqueAnswers))
	}
	fmt.Printf("Total Answers: %d\n", countAllAnswers)
}

func part2() {
	rows := readFile()
	input := strings.Join(rows, "|")
	answers := strings.Split(input, "||")

	total := 0

	for _, answer := range answers {
		answer1Line := strings.ReplaceAll(answer, "|", "")
		groupAnswers := strings.Split(answer, "|")

		uniqueAnswers := ""
		for _, letter := range answer1Line {
			if !strings.Contains(uniqueAnswers, string(letter)) {
				uniqueAnswers += string(letter)
			}
		}

		wholeGroupAnsweredYes := ""
		for _, letterR := range uniqueAnswers {
			letter := string(letterR)
			appearsInAll := true
			for _, grpAnswer := range groupAnswers {
				if !strings.Contains(grpAnswer, letter) {
					appearsInAll = false
				}
			}
			if appearsInAll {
				wholeGroupAnsweredYes += letter
			}
		}
		total += len(wholeGroupAnsweredYes)
	}
	fmt.Printf("Total Answers: %d\n", total)
}

func main() {
	part1()
	fmt.Println("\n-----------\nPART 2\n-----------")
	part2()
}
