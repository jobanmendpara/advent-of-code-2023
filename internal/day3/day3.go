package day3

import (
	"log"
	"strconv"
	"strings"

	"github.com/jobanmendpara/advent-of-code-2023/internal/helpers"
)

type number struct {
	start int
	end   int
	num   int
}

func Part1(path string) (sum int) {
	var lines []string = helpers.GetInput(path)

	for i, row := range getNumbers(lines) {
		for _, number := range row {
			if number.start != 0 && isSymbol(rune(lines[i][number.start-1])) {
				sum += number.num
				continue
			}

			if number.end < len(lines[i])-1 && isSymbol(rune(lines[i][number.end+1])) {
				sum += number.num
				continue
			}

			var k int = 0
			if number.start != 0 {
				k = number.start - 1
			}

			var end int = number.end
			if number.end < len(lines[i])-1 {
				end += 1
			}

			for ; k <= end; k++ {
				if i != 0 {
					if isSymbol(rune(lines[i-1][k])) {
						sum += number.num
						break
					}
				}

				if i != len(lines)-1 {
					if isSymbol(rune(lines[i+1][k])) {
						sum += number.num
						break
					}
				}
			}
		}
	}

	return sum
}

func Part2(path string) (sum int) {
	var lines []string = helpers.GetInput(path)

	for i, line := range lines {
		for j, char := range line {
			var count int
			var product int = 1

			if char == '*' {
				if j != 0 && helpers.IsNumber(rune(line[j-1])) {
					count++

					for _, number := range getNumbers(lines)[i] {
						if number.start <= j-1 && number.end >= j-1 {
							product *= number.num

							break
						}
					}
				}

				if j != len(line)-1 && helpers.IsNumber(rune(line[j+1])) {
					count++

					for _, num := range getNumbers(lines)[i] {
						if num.start <= j+1 && num.end >= j+1 {
							product *= num.num

							break
						}
					}
				}

				var start int
				if j != 0 {
					start = j - 1
				}

				var end int = j
				if j < len(line)-1 {
					end += 1
				}

				if i != 0 {
					for _, n := range getNumbers(lines)[i-1] {
						if (n.start <= start && n.end >= end) ||
							(n.start <= start && n.end <= end && n.end >= start) ||
							(n.start >= start && n.start <= end && n.end >= end) ||
							(n.start >= start && n.end <= end) {
							count++

							product *= n.num
						}
					}
				}

				if i != len(lines)-1 {
					for _, num := range getNumbers(lines)[i+1] {
						if (num.start <= start && num.end >= end) ||
							(num.start <= start && num.end <= end && num.end >= start) ||
							(num.start >= start && num.start <= end && num.end >= end) ||
							(num.start >= start && num.end <= end) {
							count++

							product *= num.num

						}
					}
				}
				if count == 2 {
					sum += product
				}
			}
		}
	}
	return sum
}

func getNumbers(input []string) (numbers [][]number) {
	for _, line := range input {
		var numsInLine []number = []number{}
		var strBuilder strings.Builder
		var localNumber number
		var err error

		for j, char := range line {
			if helpers.IsNumber(rune(char)) {
				strBuilder.WriteRune(char)
			}

			if !helpers.IsNumber(rune(char)) && strBuilder.Len() != 0 {
				localNumber.num, err = strconv.Atoi(strBuilder.String())
				if err != nil {
					log.Fatal(err)
				}

				localNumber.end = j - 1
				localNumber.start = localNumber.end - strBuilder.Len() + 1
				strBuilder.Reset()
				numsInLine = append(numsInLine, localNumber)
				localNumber = number{}
			}
		}

		if strBuilder.Len() != 0 {
			localNumber.num, err = strconv.Atoi(strBuilder.String())
			if err != nil {
				log.Fatal(err)
			}

			localNumber.end = len(line) - 1
			localNumber.start = localNumber.end - strBuilder.Len() + 1

			numsInLine = append(numsInLine, localNumber)
		}

		numbers = append(numbers, numsInLine)
	}

	return numbers
}

func isSymbol(r rune) bool {
	return r != '.' && !helpers.IsNumber(r)
}

func buildTable(lines []string) (table [][]string) {
	for _, line := range lines {
		var charSlice []string

		for _, char := range line {
			charSlice = append(charSlice, string(char))
		}

		table = append(table, charSlice)
	}

	return table
}
