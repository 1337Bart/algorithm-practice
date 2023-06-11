package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) []int {
	// Base case: if the list contains 0 or 1 element, it is already sorted.
	if len(arr) <= 1 {
		return arr
	}

	// Choose a pivot randomly and remove it from the list.
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]
	arr = append(arr[:pivotIndex], arr[pivotIndex+1:]...)

	smaller := make([]int, 0)
	larger := make([]int, 0)

	// Divide the remaining elements into two lists based on whether they are smaller or larger than the pivot.
	for _, num := range arr {
		if num < pivot {
			smaller = append(smaller, num)
		} else {
			larger = append(larger, num)
		}
	}

	// Recursively sort the smaller and larger lists and combine the results.
	return append(append(quickSort(smaller), pivot), quickSort(larger)...)
}

func main() {
	// Seed the random number generator to ensure we get a different output each time we run the program.
	rand.Seed(time.Now().UnixNano())

	// Generate a random list to sort.
	unsortedList := rand.Perm(1000)
	fmt.Println("Unsorted list:", unsortedList)

	// Sort the list.
	start := time.Now()
	sortedList := quickSort(unsortedList)
	duration := time.Since(start)

	fmt.Printf("Execution time: %v ms\n", duration.Milliseconds())
	fmt.Println("Sorted list:", sortedList)
}
