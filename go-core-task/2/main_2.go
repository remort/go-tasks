package main

import "fmt"

func sliceExample[T int](slice []T) []T {
	newSlice := []T{}
	for v := range slice {
		if slice[v]%2 == 0 {
			newSlice = append(newSlice, slice[v])
		}
	}
	return newSlice
}

func addElements[T int](slice []T, element T) []T {
	newSlice := make([]T, 0, len(slice)+1)
	newSlice = append(newSlice, slice...)
	newSlice = append(newSlice, element)
	return newSlice
}

func addElementsLowLevel[T int](slice []T, element T) []T {
	newSize := len(slice) + 1
	newSlice := make([]T, newSize)
	for i := range newSize {
		switch i {
		case newSize - 1:
			newSlice[i] = element
		default:
			newSlice[i] = slice[i]
		}
	}
	return newSlice
}

func copySlice[T any](slice []T) []T {
	// This function can be fully replaced with go's copy()
	// and only demonstrates how slice copying works.
	newSlice := make([]T, len(slice))
	for i := range slice {
		newSlice[i] = slice[i]
	}
	return newSlice
}

func removeElement[T any](slice []T, index int) []T {
	newSize := len(slice) - 1
	// This case can be replaced with copy() but we do things low-level here
	if index < 0 || index >= len(slice) {
		newSize = len(slice)
	}
	newSlice := make([]T, newSize)
	j := 0
	for i := range slice {
		if i != index {
			newSlice[j] = slice[i]
			j++
		}
	}
	return newSlice
}

func main() {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenNumbers := sliceExample(originalSlice)
	fmt.Println("Task 1:")
	fmt.Printf(
		"Original slice: %v. Only even numbers copy slice: %v\n",
		originalSlice, evenNumbers,
	)

	fmt.Println("\nTask 2:")
	sliceOfInts := []int{10, 20, 30}
	newElement := 40
	extendedSliceOfInts := addElementsLowLevel(sliceOfInts, newElement)
	fmt.Printf(
		"Slice of ints: %v. A copy of it with appended '%v': %v\n",
		sliceOfInts, newElement, extendedSliceOfInts,
	)

	fmt.Println("\nTask 2 (low-level version):")
	sliceOfIntsLL := []int{10, 20, 30}
	newElementLL := 40
	extendedSliceOfIntsLL := addElementsLowLevel(sliceOfIntsLL, newElementLL)
	fmt.Printf(
		"Slice of ints: %v. A copy of it with appended '%v': %v\n",
		sliceOfIntsLL, newElementLL, extendedSliceOfIntsLL,
	)

	fmt.Println("\nTask 3:")
	sliceOfChars := []byte{'r', 'o', 'a', 'd'}
	sliceOfCharsCopy := copySlice(sliceOfChars)
	fmt.Printf(
		"Original slice: %s.\nIts copy: %s\n",
		sliceOfChars, sliceOfCharsCopy,
	)
	sliceOfCharsCopy[len(sliceOfChars)-1] = 'r'
	fmt.Printf("Modified copy: %s.\nUnchanged original: %s\n", sliceOfCharsCopy, sliceOfChars)

	fmt.Println("\nTask 4:")
	products := []string{"apple", "pear", "potato"}
	fruits := removeElement(products, 2)
	fmt.Printf("A slice: %s.\nA copy of slice with last element removed: %s\n", products, fruits)
}
