package main

import (
	"fmt"
	"time"
)

func bubbleSort(unsortedList []int) []int {
	numElements := len(unsortedList)
	// Iterating over each element in the list.
	for currentIndex := 0; currentIndex < numElements; currentIndex++ {
		// For each element, iterate over the list again up to the part that remains unsorted.
		for comparisonIndex := 0; comparisonIndex < numElements-currentIndex-1; comparisonIndex++ {
			// If the current element is greater than the next one, swap them.
			if unsortedList[comparisonIndex] > unsortedList[comparisonIndex+1] {
				unsortedList[comparisonIndex], unsortedList[comparisonIndex+1] = unsortedList[comparisonIndex+1], unsortedList[comparisonIndex]
			}
		}
	}
	return unsortedList
}

func main() {
	haystack := []int{10, 12, 3, 9, 4, 6, 7, 8, 1, 10}

	start := time.Now()
	fmt.Println(bubbleSort(haystack))
	duration := time.Since(start)

	fmt.Println("Execution time:", duration)
}
