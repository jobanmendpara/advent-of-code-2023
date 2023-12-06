package day4

import (
	"fmt"
	_ "fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jobanmendpara/advent-of-code-2023/internal/helpers"
)

type game struct {
	id      int
	numbers []int
	winners []int
}

func Part1(inputFilePath string) (sum int) {
	var lines []string = helpers.GetInput(inputFilePath)

	for _, game := range lines {
		var points int
		var myNumbers []int
		var winningNumbers []int
		var _, newString string = helpers.GetStringInBetween(game, "Card ", ":")
		var isFirstMatch bool = true

		myNumbers = convertArrayStringToArrayInt(strings.Split(strings.Split(newString, " | ")[0], " "))
		winningNumbers = convertArrayStringToArrayInt(strings.Split(strings.Split(newString, " | ")[1], " "))

		for _, num := range myNumbers {
			for _, winner := range winningNumbers {
				if num == winner {
					if isFirstMatch {
						points += 1
						isFirstMatch = false
					} else {
						points *= 2
					}
				}
			}
		}

		sum += points
	}

	return sum
}

// WARN: Extremely inefficient!
func Part2(inputFilePath string) (cards int) {
	var lines []string = helpers.GetInput(inputFilePath)
	var table map[int][]game = buildTable(lines)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(table[i]); j++ {
			var game = table[i][j]
			var matches int

			for _, num := range game.numbers {
				for _, winner := range game.winners {
					if num == winner {
						matches++
					}
				}
			}

			for k := 0; k < matches; k++ {
				table[i+k+1] = append(table[i+k+1], table[i+k+1][0])
			}
		}
	}

	for _, games := range table {
		for range games {
			cards++
		}
	}

	return cards
}

func buildTable(input []string) (table map[int][]game) {
	table = make(map[int][]game)
	for i, row := range input {
		var gameIdString, newString string = helpers.GetStringInBetween(row, "Card ", ":")
		var numbers []int = convertArrayStringToArrayInt(strings.Split(strings.Split(newString, " | ")[0], " "))
		var winners []int = convertArrayStringToArrayInt(strings.Split(strings.Split(newString, " | ")[1], " "))

		var id, err = strconv.Atoi(strings.TrimSpace(gameIdString))
		if err != nil {
			fmt.Println(i)
			log.Fatal(err)
		}

		var game game = game{
			id:      id,
			numbers: numbers,
			winners: winners,
		}

		table[i] = append(table[i], game)
	}

	return table
}

func convertArrayStringToArrayInt(input []string) (array []int) {
	for _, str := range input {
		if len(str) > 0 {
			number, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}

			array = append(array, number)
		}
	}

	return array
}
