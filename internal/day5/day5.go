package day5

import (
	"strings"

	"github.com/jobanmendpara/advent-of-code-2023/internal/helpers"
)

type table map[uint64][]uint64

func Part1(inputFilePath string) (lowestLocation uint64) {
	var lines []string = helpers.GetInput(inputFilePath)
	var locations []uint64

	seeds := helpers.ConvertArrayStringToArrayUint(strings.Split(lines[0][strings.Index(lines[0], ":")+1:], " "))
	lines = lines[2:]

	startingMap, remaining := getNextMap(lines)

	for _, seed := range seeds {
		location := travelMaps(uint64(seed), startingMap, remaining)

		locations = append(locations, location)
	}

	lowestLocation = helpers.FindMinimum(locations)

	return lowestLocation
}

func Part2(inputFilePath string) (lowestLocation uint64) {
	var _ []string = helpers.GetInput(inputFilePath)

	return lowestLocation
}

func getNextMap(lines []string) (res table, newLines []string) {
	res = make(table)
	var count uint64

	for i, line := range lines {
		count++

		if uint64(len(line)) < 1 {
			newLines = lines[count:]
			return res, newLines
		} else if !helpers.IsNumber(rune(line[0])) {
			continue
		} else {
			res[uint64(i)] = helpers.ConvertArrayStringToArrayUint(strings.Split(line, " "))
		}
	}

	return res, newLines
}

func findLocation(input uint64, table table) (location uint64) {
	for i, row := range table {
		dest := row[0]
		src := row[1]
		rng := row[2]

		if input >= src && input < src+rng {
			location = dest + (input - src)

			return location
		} else if i == uint64(len(table)) {
			location = input

			return location
		}
	}

	return location
}

func travelMaps(start uint64, table table, rows []string) (end uint64) {
	if len(rows) < 1 {
		end = start
	} else {
		nextStart := findLocation(start, table)
		nextMap, remaining := getNextMap(rows)
		end = travelMaps(nextStart, nextMap, remaining)
	}

	return end
}
