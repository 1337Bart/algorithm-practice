package main

import (
	"fmt"
	"time"
)

func linearSearch(listToSearch []int, targetValue int) bool {
	// Iterate over each element in the list.
	for currentIndex := 0; currentIndex < len(listToSearch); currentIndex++ {
		// If the current element is equal to the target value, the search is successful.
		if listToSearch[currentIndex] == targetValue {
			return true // Target found.
		}
	}
	return false // Target not found.
}

func main() {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	needle := 5

	start := time.Now()
	fmt.Println(linearSearch(haystack, needle))
	duration := time.Since(start)

	fmt.Println("Execution time:", duration)
}
