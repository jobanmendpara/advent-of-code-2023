package day1

import (
  "strconv"
  "bufio"
  "log"
  "path/filepath"
  "os"
  "regexp"
)

func Solve() int {
  var sum int = 0

  path, err := filepath.Abs("internal/day1/input.txt")
  if err != nil {
    log.Fatal(err)
  }

  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    justNums, err := stripNonNumeric(string(scanner.Text()))
    if err != nil {
      log.Fatal(err)
    }

    runes := []rune(justNums)
    firstNum := runes[0]
    lastNum := runes[len(runes)-1]

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

  return sum
}

func stripNonNumeric(str string) (string, error) {
  reg, err := regexp.Compile("[^0-9]+")
  if err != nil {
    log.Fatal(err)
  }

  return reg.ReplaceAllString(str, ""), nil
}
