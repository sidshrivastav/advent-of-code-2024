package solution

import (
	"advent-of-code-2024/pkg/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func TotalDistance() (int, error) {
	location_id_1, location_id_2, shouldReturn, _, err := getData()
	if shouldReturn {
		return 0, err
	}

	sort.Ints(location_id_1)
	sort.Ints(location_id_2)

	var total_distance int
	for index := range location_id_1 {
		// Calculate the absolute difference between corresponding values
		total_distance += common.Max(location_id_1[index], location_id_2[index]) - common.Min(location_id_1[index], location_id_2[index])
	}

	return total_distance, nil
}

func SimilarityScore() (int, error) {
	location_id_1, location_id_2, shouldReturn, _, err := getData()
	if shouldReturn {
		return 0, err
	}

	countMap := common.CountOccurrences(location_id_2)

	var score int
	for _, value := range location_id_1 {
		// Calculate the absolute difference between corresponding values
		if count, exists := countMap[value]; exists {
			score += value * count
		}

	}

	return score, nil
}

func getData() ([]int, []int, bool, int, error) {
	/*
		Input:
		3   4
		4   3
		2   5
		1   3
		3   9
		3   3
	*/
	content, err := common.GetContent("input.txt")
	if err != nil {
		return nil, nil, true, -1, fmt.Errorf("error opening file: %w", err)
	}

	var location_id_1 []int
	var location_id_2 []int

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		values := strings.Fields(line)
		num1, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, nil, true, -1, fmt.Errorf("error converting value to int: %w", err)

		}
		location_id_1 = append(location_id_1, num1)

		num2, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, nil, true, -1, fmt.Errorf("error converting value to int: %w", err)

		}
		location_id_2 = append(location_id_2, num2)
	}
	return location_id_1, location_id_2, false, 0, nil
}
