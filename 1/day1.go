package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() []int {
	var inputs []int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, e := strconv.Atoi(scanner.Text())
		if e != nil {
			log.Fatal(e)
		}
		inputs = append(inputs, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func part1() {
	inputs := readFile()

	for i, val := range inputs[0 : len(inputs)-2] {
		for _, nextVal := range inputs[i+1 : len(inputs)-1] {
			if val+nextVal == 2020 {
				fmt.Printf("\n\ni: %d, j: %d, sum: %d, product: %d\n", val, nextVal, val+nextVal, val*nextVal)
			}
		}
	}

}

func part2() {
	inputs := readFile()

	for i, val := range inputs[0 : len(inputs)-2] {
		for j, nextVal := range inputs[i+1 : len(inputs)-1] {
			for _, thirdVal := range inputs[j+1 : len(inputs)-2] {
				if val+nextVal+thirdVal == 2020 {
					fmt.Printf("\n\ni: %d, j: %d, k: %d, sum: %d, product: %d\n", val, nextVal, thirdVal, val+nextVal+thirdVal, val*nextVal*thirdVal)
					return
				}
			}
		}
	}
}

func main() {
	part1()
	fmt.Println("\n-----------\nPART 2\n-----------")
	part2()
}
