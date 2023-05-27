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

func binarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := (low + high) / 2
		if haystack[mid] < needle {
			low = mid + 1
		} else if haystack[mid] > needle {
			high = mid - 1
		} else {
			return true // needle found
		}
	}
	return false // needle not found
}

func main() {
	haystack := generateSortedArray(1000000)

	//select a needle from the haystack pool of numbers
	needle := haystack[rand.Intn(len(haystack))]

	start := time.Now()
	fmt.Println(binarySearch(haystack, needle))
	duration := time.Since(start)

	fmt.Println("Execution time:", duration)
}
