package main

import (
	"advent-of-code-2024/day01/solution"
	"fmt"
)

func main() {
	result, err := solution.SimilarityScore()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	fmt.Println(result)
}
