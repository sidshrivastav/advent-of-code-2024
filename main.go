package main

import (
	"advent-of-code-2024/day02/solution"
	"fmt"
)

func main() {
	result, err := solution.SafeReportsWithDampener()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	fmt.Println(result)
}
