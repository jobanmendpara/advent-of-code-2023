package helpers

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func IsNumber(r rune) bool {
  return r >= '0' && r <= '9'
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
