package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func part1() ([]int, int) {
	rows := readFile()
	maxSeatID := 0
	var allSeatIDs []int
	for _, boardingPass := range rows {
		rowCode := boardingPass[0:7]
		columnCode := boardingPass[7:10]
		row := 0
		column := 0
		for i, rowLetter := range rowCode {
			if string(rowLetter) == "B" {
				row += int(math.Exp2(float64(6 - i)))
			}
		}

		for i, columnLetter := range columnCode {
			if string(columnLetter) == "R" {
				column += int(math.Exp2(float64(2 - i)))
			}
		}

		seatID := row*8 + column
		allSeatIDs = append(allSeatIDs, seatID)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}

		// fmt.Printf("rowCode: %s, columnCode: %s, row: %d, column: %d, seatID: %d\n", rowCode, columnCode, row, column, seatID)

	}

	fmt.Printf("MAX SEAT ID: %d\n", maxSeatID)
	return allSeatIDs, maxSeatID
}

func part2() {
	allSeatIDs, _ := part1()

	var missingIDs []int

	sort.Ints(allSeatIDs)

	for i := range allSeatIDs {
		if (i > 0) && (i < len(allSeatIDs)-1) {
			if allSeatIDs[i]-1 != allSeatIDs[i-1] || allSeatIDs[i]+1 != allSeatIDs[i+1] {
				missingIDs = append(missingIDs, allSeatIDs[i])
			}
		}
	}

	if len(missingIDs) == 2 {
		fmt.Printf("My Seat: %d\n", missingIDs[0]+1)
		fmt.Printf("Just Checking: %d\n", missingIDs[1]-1)
	}

}

func main() {
	part1()
	fmt.Println("\n-----------\nPART 2\n-----------")
	part2()
}
