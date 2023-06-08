package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func generateSortedArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	sort.Ints(arr) // Ensure array is sorted
	return arr
}

func binarySearch(sortedList []int, targetValue int) bool {
	low := 0
	high := len(sortedList) - 1

	// While the list has not been fully traversed.
	for low <= high {
		mid := (low + high) / 2
		// If the middle value is less than the target, discard the lower half of the list.
		if sortedList[mid] < targetValue {
			low = mid + 1
		} else if sortedList[mid] > targetValue {
			// If the middle value is more than the target, discard the upper half of the list.
			high = mid - 1
		} else {
			return true // Target found.
		}
	}
	return false // Target not found.
}

func main() {
	haystack := generateSortedArray(1000000)

	// Select a needle from the haystack pool of numbers
	needle := haystack[rand.Intn(len(haystack))]

	start := time.Now()
	found := (binarySearch(haystack, needle))

	if found == true {
		fmt.Println("Needle found.")
	}

	duration := time.Since(start)

	fmt.Println("Execution time:", duration)
}
