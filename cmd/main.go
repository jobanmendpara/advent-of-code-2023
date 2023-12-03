package main

import (
	"fmt"
	"github.com/jobanmendpara/aoc23/internal/day2"
)

func main() {
	fmt.Println("---START---")

	sum, poweredSum := day2.Solve()

	fmt.Println("Sum:", sum)
	fmt.Println("PoweredSum:", poweredSum)
	fmt.Println("---END---")
}
