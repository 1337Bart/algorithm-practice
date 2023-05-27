package main

import "fmt"

func bubbleSort(haystack []int) []int {
	n := len(haystack)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if haystack[j] > haystack[j+1] {
				haystack[j], haystack[j+1] = haystack[j+1], haystack[j]
			}
		}
	}
	return haystack
}

func main() {
	haystack := []int{10, 12, 3, 9, 4, 6, 7, 8, 1, 10}

	fmt.Println(bubbleSort(haystack))
}
