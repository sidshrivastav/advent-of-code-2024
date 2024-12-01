package common

func CountOccurrences(arr []int) map[int]int {
	// Create a map to store the count of each element
	countMap := make(map[int]int)

	// Iterate through the array
	for _, num := range arr {
		// Increment the count for the number in the map
		countMap[num]++
	}

	// Return the count map
	return countMap
}
