package main

import (
	"fmt"
	"time"
)

func insertionSort(unsortedList []int) []int {
	// We start from the second element since the first one is already sorted by itself.
	for currentIndex := 1; currentIndex < len(unsortedList); currentIndex++ {
		// This is the value we want to insert into the sorted part of the list.
		valueToInsert := unsortedList[currentIndex]
		// We start comparing with the element before it.
		comparisonIndex := currentIndex - 1

		// While the comparison index is valid (not out of range)
		// and the value at this index is greater than the value we want to insert,
		// we shift the larger values to the right.
		for comparisonIndex >= 0 && unsortedList[comparisonIndex] > valueToInsert {
			unsortedList[comparisonIndex+1] = unsortedList[comparisonIndex]
			comparisonIndex--
		}

		// Now we have found the correct position for the value,
		// we insert it into its proper place in the sorted part of the list.
		unsortedList[comparisonIndex+1] = valueToInsert
	}
	return unsortedList
}

func main() {
	haystack := []int{10, 12, 3, 9, 4, 6, 7, 8, 1, 10}

	start := time.Now()
	fmt.Println(insertionSort(haystack))
	duration := time.Since(start)

	fmt.Println("Execution time:", duration)
}
