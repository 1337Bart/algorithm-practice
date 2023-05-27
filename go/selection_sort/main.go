package main

import "fmt"

func selectionSort(haystack []int) []int {
	for i := 0; i < len(haystack); i++ {
		minIdx := i
		for j := i + 1; j < len(haystack); j++ {
			if haystack[j] < haystack[minIdx] {
				minIdx = j
			}
		}
		haystack[i], haystack[minIdx] = haystack[minIdx], haystack[i]
	}
	return haystack
}

func main() {
	haystack := []int{10, 12, 3, 9, 4, 6, 7, 8, 1, 10}

	fmt.Println(selectionSort(haystack))
}
