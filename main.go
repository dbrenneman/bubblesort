package main

import (
	"fmt"
	"math/rand"
)

func makeRandomSlice(numItems, max int) []int {
	s := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}

	return s
}

func printSlice(s []int, max int) {
	for i, v := range s {
		if i < max {
			fmt.Println(v)
		} else {
			break
		}
	}
}

func bubbleSort(s []int) {
	changed := true
	for changed {
		changed = false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				s[i-1], s[i] = s[i], s[i-1]
				changed = true
			}
		}
	}
}

func checkSorted(s []int) {
	sorted := true
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			sorted = false
		}
	}

	if sorted {
		fmt.Println("The slice is sorted.")
	} else {
		fmt.Println("The slice is NOT sorted!")
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println("pre-sorted ^^^")

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
