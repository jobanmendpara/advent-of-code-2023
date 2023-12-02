package helpers

import (
	"log"
	"os"
	"path/filepath"
)

func IsNumber(r rune) bool {
  return r >= '0' && r <= '9'
}

func GetInput(inputFilePath string) *os.File {
	path, err := filepath.Abs(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
