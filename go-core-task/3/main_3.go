package main

import (
	"fmt"
	"maps"
)

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() *StringIntMap {
	return &StringIntMap{
		data: maps.Clone(m.data),
	}
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}

func main() {
	mapping := NewStringIntMap()
	fmt.Printf("New mapping: %v\n", mapping)

	mapping.Add("year", 2026)
	mapping.Add("month", 3)
	mapping.Add("day", 1)
	mapping.Remove("month")
	mapping.Add("month", 4)
	month, _ := mapping.Get("month")
	fmt.Printf("The 'month' stored in the mapping: %d\n", month)

	fmt.Printf("Mapping holds current date: %v\n", mapping)

	fmt.Printf("Mapping holds 'day': %v\n", mapping.Exists("day"))
	fmt.Printf("Mapping holds 'hour': %v\n", mapping.Exists("hour"))

	newMapping := mapping.Copy()
	newMapping.Add("hour", 18)
	newMapping.Add("minute", 34)
	newMapping.Add("second", 59)
	fmt.Printf("First mapping: %v\n", mapping)
	fmt.Printf("Second mapping: %v\n", newMapping)
}
