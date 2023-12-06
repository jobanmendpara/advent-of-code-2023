package helpers

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func IsNumber(r rune) bool {
  return r >= '0' && r <= '9'
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

  for (scanner.Scan()) {
    input = append(input, scanner.Text())
  }

	return input
}
