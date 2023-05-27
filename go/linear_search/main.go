package main

import "fmt"

func linearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true // needle found
		}
	}
	return false // needle not found
}

func main() {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	needle := 5
	fmt.Println(linearSearch(haystack, needle))
}
