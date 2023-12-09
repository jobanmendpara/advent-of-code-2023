package main

import (
	"fmt"

	"github.com/jobanmendpara/advent-of-code-2023/internal/day5"
)

func main() {
	fmt.Println("---START---")

  fmt.Println("Example - Pt1:", day5.Part1("internal/day5/example.txt"))
  fmt.Println("Actual - Pt1:", day5.Part1("internal/day5/input.txt"))
	fmt.Println("Example - Pt2:", day5.Part2("internal/day5/example.txt"))
	fmt.Println("Input - Pt2:", day5.Part2("internal/day5/input.txt"))

  fmt.Println("---END---")
}
