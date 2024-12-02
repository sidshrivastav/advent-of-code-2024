package solution

import (
	"advent-of-code-2024/pkg/common"
	"fmt"
	"strconv"
	"strings"
)

func isSafeReport(report []int) bool {
	if len(report) < 2 {
		return false // A single level can't be "increasing" or "decreasing"
	}

	// Determine the trend: increasing or decreasing
	increasing := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		// Check if adjacent levels differ by at least 1 and at most 3
		if diff < 1 && diff > -1 || diff > 3 || diff < -3 {
			return false
		}

		// Check if the trend remains consistent
		if (increasing && report[i+1] <= report[i]) || (!increasing && report[i+1] >= report[i]) {
			return false
		}
	}

	return true
}

func isSafeReportWithDampener(report []int) bool {
	if isSafeReport(report) {
		return true
	}
	fmt.Println(report)
	for index := range report {
		firstPart := report[:index] // Elements from the start to the index (exclusive)
		secondPart := report[index+1:]
		firstPartCopy := make([]int, len(firstPart))
		secondPartCopy := make([]int, len(secondPart))
		copy(firstPartCopy, firstPart)
		copy(secondPartCopy, secondPart)
		tempReport := append(firstPartCopy, secondPartCopy...)
		fmt.Println(index, report, tempReport)
		if isSafeReport(tempReport) {
			return true
		}
	}
	return false
}

func SafeReports() (int, error) {
	safeReport := 0
	reports, shouldReturn, _, err := getData()
	if shouldReturn {
		return 0, err
	}
	for _, report := range reports {
		if isSafeReport(report) {
			safeReport++
		}
	}
	return safeReport, nil
}

func SafeReportsWithDampener() (int, error) {
	safeReport := 0
	reports, shouldReturn, _, err := getData()
	if shouldReturn {
		return 0, err
	}
	for _, report := range reports {
		if isSafeReportWithDampener(report) {
			safeReport++
		}
	}
	return safeReport, nil
}

func getData() ([][]int, bool, int, error) {
	/*
		Input:
		7 6 4 2 1
		1 2 7 8 9
		9 7 6 2 1
		1 3 2 4 5
		8 6 4 4 1
		1 3 6 7 9
	*/
	content, err := common.GetContent("input.txt")
	if err != nil {
		return nil, true, -1, fmt.Errorf("error opening file: %w", err)
	}

	var report [][]int

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		values := strings.Fields(line)
		var level []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, true, -1, fmt.Errorf("error converting value to int: %w", err)
			}
			level = append(level, num)
		}

		report = append(report, level)
	}
	return report, false, 0, nil
}
