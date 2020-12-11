package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	rows := readFile()
	rowsWidth := len(rows[0])

	y := 0
	treesEncountered := 0

	for _, row := range rows {
		if string(row[y]) == "#" {
			treesEncountered++
		}
		y = (y + 3) % rowsWidth
	}
	fmt.Printf("Trees Encountered: %d\n", treesEncountered)
}

func checkPath(right int, down int) int {
	rows := readFile()
	rowsWidth := len(rows[0])

	y := 0
	treesEncountered := 0

	for i, row := range rows {
		if i%down == 0 {
			if string(row[y]) == "#" {
				treesEncountered++
			}
			y = (y + right) % rowsWidth
		}
	}
	return treesEncountered
}

func part2() {

	path1 := checkPath(1, 1)
	path2 := checkPath(3, 1)
	path3 := checkPath(5, 1)
	path4 := checkPath(7, 1)
	path5 := checkPath(1, 2)

	fmt.Printf("Path1: %d, Path2: %d, Path3: %d, Path4: %d, Path5: %d\n\n", path1, path2, path3, path4, path5)

	product := path1 * path2 * path3 * path4 * path5
	fmt.Printf("Product: %d\n", product)
}

func main() {
	part1()
	fmt.Println("\n-----------\nPART 2\n-----------")
	part2()
}
