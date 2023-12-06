package main

import (
	"fmt"

	"github.com/jobanmendpara/advent-of-code-2023/internal/day4"
)

func main() {
  var input string = "internal/day4/input.txt"

	fmt.Println("---START---")

	fmt.Println("Part1:", day4.Part1(input))
	fmt.Println("Part2:", day4.Part2(input))

  fmt.Println("---END---")
}
