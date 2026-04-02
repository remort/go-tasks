package main

import (
	"fmt"
	"slices"
)

func intersectSlice(s1 []int, s2 []int) (bool, []int) {
	slices.Sort(s1)
	slices.Sort(s2)
	fmt.Printf("sorted1: %v, sorted2: %v\n", s1, s2)
	tempSlice := make([]int, len(s1))
	m := 0
	k := 0
	found := false
	prev := 0
	for i1, v1 := range s1 {
		if i1 > 0 && found && v1 == prev {
			prev = v1
			continue
		}
		found = false
		if slices.Contains(s2[m:], v1) {
			tempSlice[k] = v1
			k++
			m++
			found = true

		}
		prev = v1
	}
	resultingSlice := make([]int, k)
	copy(resultingSlice, tempSlice)
	fmt.Printf("S1: %v, S2: %+v\n", tempSlice, resultingSlice)
	return len(resultingSlice) > 0, resultingSlice
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	ok, intersection := intersectSlice(a, b)
	fmt.Printf("Slices intersection: %v, %v\n", ok, intersection)

	x := []int{65, 3, 42, 678, 64}
	y := []int{42, 42, 2, 3, 43}
	ok, intersection = intersectSlice(x, y)
	fmt.Printf("Slices intersection: %v, %v\n", ok, intersection)
}
