package helpers

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ConvertArrayStringToArrayUint(input []string) (array []uint64) {
	for _, str := range input {
		if len(str) > 0 {
			number, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			array = append(array, number)
		}
	}

	return array
}

func GetStringInBetween(input string, start string, end string) (result string, newString string) {
	var startIndex = strings.Index(input, start) + len(start)
	if startIndex == -1 {
		return
	}

	var endIndex = strings.Index(input, end) + len(end) - 1
	if endIndex == -1 {
		return
	}

	result = input[startIndex:endIndex]
	newString = input[endIndex+2:]

	return result, newString
}

func GetInput(inputFilePath string) (input []string) {
	path, err := filepath.Abs(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func FindMinimum(arr []uint64) (min uint64) {
	if len(arr) == 0 {
		return 0
	}

	min = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}

  return min
}

func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
