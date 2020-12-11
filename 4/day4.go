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
	rows := readFile()
	input := strings.Join(rows, "|")
	passports := strings.Split(input, "||")

	validPassports := 0

	for _, passport := range passports {

		passport1Line := strings.ReplaceAll(passport, "|", " ")
		passportEntries := strings.Split(passport1Line, " ")
		validEntries := 0

		for _, passportEntry := range passportEntries {
			key := strings.Split(passportEntry, ":")

			switch key[0] {
			case "byr":
				validEntries++
			case "iyr":
				validEntries++
			case "eyr":
				validEntries++
			case "hgt":
				validEntries++
			case "hcl":
				validEntries++
			case "ecl":
				validEntries++
			case "pid":
				validEntries++
			}
		}

		if validEntries == 7 {
			validPassports++
		}
	}
	fmt.Printf("# Valid Passports: %d\n", validPassports)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func part2() {
	rows := readFile()
	input := strings.Join(rows, "|")
	passports := strings.Split(input, "||")

	validPassports := 0

	for _, passport := range passports {

		passport1Line := strings.ReplaceAll(passport, "|", " ")
		passportEntries := strings.Split(passport1Line, " ")
		validEntries := 0
		// fmt.Println(passport1Line)
		for _, passportEntry := range passportEntries {

			key := strings.Split(passportEntry, ":")

			switch key[0] {
			case "byr":
				byr, byrE := strconv.Atoi(key[1])
				checkErr(byrE)
				if byr >= 1920 && byr <= 2002 {
					validEntries++
				}
			case "iyr":
				iyr, iyrE := strconv.Atoi(key[1])
				checkErr(iyrE)
				if iyr >= 2010 && iyr <= 2020 {
					validEntries++
				}
			case "eyr":
				eyr, eyrE := strconv.Atoi(key[1])
				checkErr(eyrE)
				if eyr >= 2020 && eyr <= 2030 {
					validEntries++
				}
			case "hgt":
				units := key[1][len(key[1])-2:]
				value, _ := strconv.Atoi(strings.TrimSuffix(key[1], units))
				if units == "cm" && value >= 150 && value <= 193 {
					validEntries++
					// fmt.Printf("hgt OK: %s\n", key[1])
				}

				if units == "in" && value >= 59 && value <= 76 {
					validEntries++
					// fmt.Printf("hgt OK: %s\n", key[1])
				}
			case "hcl":
				r := regexp.MustCompile("^#[0-9a-f]{6}$")
				if r.MatchString(key[1]) {
					validEntries++
				}

			case "ecl":
				switch key[1] {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					validEntries++
					// fmt.Printf("ecl OK: %s\n", key[1])
				}
			case "pid":
				r := regexp.MustCompile("^[0-9]{9}$")
				if r.MatchString(key[1]) {
					validEntries++
					// fmt.Printf("pid OK: %s\n", key[1])
				}
			}
		}

		if validEntries == 7 {
			// fmt.Println("VALID")
			validPassports++
		}
	}
	fmt.Printf("# Valid Passports: %d\n", validPassports)
}

func main() {
	part1()
	fmt.Println("\n-----------\nPART 2\n-----------")
	part2()
}
