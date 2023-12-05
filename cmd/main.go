package main

import (
	"fmt"

	"github.com/jobanmendpara/advent-of-code-2023/internal/day3"
)

func main() {
	fmt.Println("---START---")

  var part1 int = day3.Solve1("internal/day3/input.txt")
  var part2 int = day3.Solve2("internal/day3/input.txt")

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
	fmt.Println("---END---")
}
