package day2

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/jobanmendpara/advent-of-code-2023/internal/helpers"
)

func Solve() (sum int, poweredSum float64) {
	const redMax int = 12
	const greenMax int = 13
	const blueMax int = 14

	var lines []string = helpers.GetInput("internal/day2/input.txt")

	for _, game := range lines {
		var redMinRequired float64 = 0
		var greenMinRequired float64 = 0
		var blueMinRequired float64 = 0
		var gameId string
		var rounds []string
		var newString string
		var isGamePossible bool = true
		var gamePower float64 = 1

		gameId, newString = helpers.GetStringInBetween(game, "Game: ", ":")
		rounds = strings.Split(newString, "; ")

		for i := range rounds {
			var splitByColor []string = strings.Split(rounds[i], ", ")

			for _, val := range splitByColor {
				var split []string = strings.Split(val, " ")
				convertedNum, err := strconv.Atoi(split[0])
				if err != nil {
					log.Fatal(err)
				}
				switch split[1] {
				case "red":
					if convertedNum > redMax {
						isGamePossible = false
					}

					redMinRequired = math.Max(redMinRequired, float64(convertedNum))
				case "blue":
					if convertedNum > blueMax {
						isGamePossible = false
					}

					blueMinRequired = math.Max(blueMinRequired, float64(convertedNum))
				case "green":
					if convertedNum > greenMax {
						isGamePossible = false
					}

					greenMinRequired = math.Max(greenMinRequired, float64(convertedNum))
				default:
					isGamePossible = false
				}
			}
		}

		if isGamePossible {
			convertedGameId, err := strconv.Atoi(gameId)
			if err != nil {
				log.Fatal(err)
			}

			sum += convertedGameId
		}

		gamePower *= (redMinRequired * blueMinRequired * greenMinRequired)
		poweredSum += gamePower
	}

	return sum, poweredSum
}
