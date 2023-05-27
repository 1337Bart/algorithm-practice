package main

import "fmt"

func insertionSort(haystack []int) []int {
	for i := 1; i < len(haystack); i++ {
		needle := haystack[i]
		j := i - 1
		for j >= 0 && haystack[j] > needle {
			haystack[j+1] = haystack[j]
			j--
		}
		haystack[j+1] = needle
	}
	return haystack
}

func main() {
	haystack := []int{10, 12, 3, 9, 4, 6, 7, 8, 1, 10}

	fmt.Println(insertionSort(haystack))
}
