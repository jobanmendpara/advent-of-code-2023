package day1

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/jobanmendpara/advent-of-code-2023/internal/helpers"
)

func Solve1() int {
	var file *os.File = helpers.GetInput("internal/day1/input.txt")

	var sum int = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		justNums, err := stripNonNumeric(string(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}

		split := []rune(justNums)
		firstNum := split[0]
		lastNum := split[len(split)-1]

		concatenated := string(firstNum) + string(lastNum)

		calibrationVal, err := strconv.Atoi(concatenated)
		if err != nil {
			log.Fatal(err)
		}

		sum += calibrationVal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return sum
}

func Solve2() int {
	var numWordMap = map[string]int{
		"zero":  0,
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

	var sum int = 0
	var file *os.File = helpers.GetInput("internal/day1/input.txt")
	var scanner *bufio.Scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		var line string = string(scanner.Text())
		var numbersInLine string = ""

		for i, rune := range line {
			if helpers.IsNumber(rune) {
				numbersInLine += string(rune)
			} else {
				for word, num := range numWordMap {
					if i+len(word) > len(line) {
						continue
					}

					var stringToCheck string = line[i : i+len(word)]

					if stringToCheck == word {
						numbersInLine += strconv.Itoa(num)
					}
				}
			}
		}

		var split []rune = []rune(numbersInLine)
		firstNum := numbersInLine[0]
		lastNum := numbersInLine[len(split)-1]

		concatenated := string(firstNum) + string(lastNum)

		calibrationVal, err := strconv.Atoi(concatenated)
		if err != nil {
			log.Fatal(err)
		}

		sum += calibrationVal
	}

	defer file.Close()

	return sum
}

func stripNonNumeric(str string) (string, error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(str, ""), nil
}
