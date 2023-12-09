package day5

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/jobanmendpara/advent-of-code-2023/internal/helpers"
)

type MapFilter struct {
	filters []filter
	chain   *MapFilter
}

type filter struct {
	start int
	end   int
	diff  int
}

func Part1(path string) int {
	input := helpers.GetInput(path)
	mapFilters := buildFilters(input)
	smallestNumber := int(^uint(0) >> 1)
	startFilter := mapFilters[0]
	seeds := strings.Fields(strings.Split(input[0], ":")[1])

	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		result := startFilter.Calculate(seedInt)
		if result < smallestNumber {
			smallestNumber = result
		}
	}

	return smallestNumber
}

func Part2(path string) int {
	input := helpers.GetInput(path)
	mapFilters := buildFilters(input)
	smallestNumber := int(^uint(0) >> 1)
	startFilter := mapFilters[0]
	seeds := buildSeeds(strings.Fields(strings.Split(input[0], ":")[1]))

	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		result := startFilter.Calculate(seedInt)
		if result < smallestNumber {
			smallestNumber = result
		}
	}



	return smallestNumber
}

func (m *MapFilter) Calculate(input int) (output int) {
	output = input

	for _, filter := range m.filters {
		if input >= filter.start && input <= filter.end {
			output = input + filter.diff
			break
		}
	}
	if m.chain != nil {
		output = m.chain.Calculate(output)
	}

	return output
}

func buildFilters(input []string) (filters []*MapFilter) {
	for i := 2; i < len(input)-1; i++ {
		if len(input[i]) == 0 || input[i] == "" {
			continue
		}
		if unicode.IsDigit(rune(input[i][0])) {
			values := strings.Fields(input[i])
			inputStart, _ := strconv.Atoi(values[1])
			filterRange, _ := strconv.Atoi(values[2])
			outputStart, _ := strconv.Atoi(values[0])

			currentFilter := filters[len(filters)-1]
			currentFilter.filters = append(currentFilter.filters, filter{
				start: inputStart,
				end:   inputStart + (filterRange - 1),
				diff:  outputStart - inputStart,
			})
		} else {
			newFilter := &MapFilter{}
			if len(filters) > 0 {
				filters[len(filters)-1].chain = newFilter
			}
			filters = append(filters, newFilter)
		}
	}

	return filters
}

func buildSeeds(input []string) (newSeeds []string) {
	for i := 0; i < len(input); i += 2 {
		seed, _ := strconv.Atoi(input[i])
		diff, _ := strconv.Atoi(input[i+1])

		for j := seed; j < seed+diff; j++ {
			newSeeds = append(newSeeds, strconv.Itoa(j))
		}
	}
	return newSeeds
}
