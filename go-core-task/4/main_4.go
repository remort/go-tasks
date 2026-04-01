package main

import (
	"fmt"
	"slices"
)

func diffSlice(s1 []string, s2 []string) []string {
	slices.Sort(s1)
	slices.Sort(s2)
	fmt.Printf("sorted1: %s, sorted2: %s\n", s1, s2)
	tempSlice := make([]string, len(s1))
	m := 0
	k := 0
	found := false
	prev := ""
	for i1, v1 := range s1 {
		if i1 > 0 && found && v1 == prev {
			prev = v1
			continue
		}
		found = false
		if slices.Contains(s2[m:], v1) {
			m++
			found = true
		}
		if !found {
			tempSlice[k] = v1
			k++
		}
		prev = v1
	}
	resultingSlice := make([]string, k)
	copy(resultingSlice, tempSlice)
	fmt.Printf("S1: %q, S2: %+v\n", tempSlice, resultingSlice)
	return resultingSlice
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	slice3 := diffSlice(slice1, slice2)
	fmt.Printf("diff slice: %v\n", slice3)

	slice4 := []string{"banana", "date", "fig"}
	slice5 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice6 := diffSlice(slice4, slice5)
	fmt.Printf("diff slice: %v\n", slice6)

	slice7 := []string{"apple", "banana", "cherry"}
	slice8 := []string{"banana"}
	slice9 := diffSlice(slice7, slice8)
	fmt.Printf("diff slice: %v\n", slice9)

}
