package main

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func TestDiffSliceBasic(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want []string
	}{
		{
			name: "s1 longer than s2",
			s1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			s2:   []string{"banana", "date", "fig"},
			want: []string{"43", "apple", "cherry", "gno1", "lead"},
		},
		{
			name: "s2 longer than s1",
			s1:   []string{"apple", "banana"},
			s2:   []string{"apple", "banana", "cherry", "date"},
			want: []string{},
		},
		{
			name: "equal slices",
			s1:   []string{"a", "b", "c"},
			s2:   []string{"a", "b", "c"},
			want: []string{},
		},
		{
			name: "no common elements",
			s1:   []string{"a", "b", "c"},
			s2:   []string{"d", "e", "f"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "empty s1",
			s1:   []string{},
			s2:   []string{"a", "b", "c"},
			want: []string{},
		},
		{
			name: "empty s2",
			s1:   []string{"a", "b", "c"},
			s2:   []string{},
			want: []string{"a", "b", "c"},
		},
		{
			name: "both empty",
			s1:   []string{},
			s2:   []string{},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffSlice(tt.s1, tt.s2)
			if !slices.Equal(got, tt.want) {
				t.Errorf("diffSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffSliceDuplicates(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want []string
	}{
		{
			name: "duplicates in s1",
			s1:   []string{"a", "b", "a", "c", "b"},
			s2:   []string{"b"},
			want: []string{"a", "a", "c"},
		},
		{
			name: "duplicates in s2",
			s1:   []string{"a", "b", "c"},
			s2:   []string{"b", "b", "b"},
			want: []string{"a", "c"},
		},
		{
			name: "unique duplicates in s2",
			s1:   []string{"a", "b", "c"},
			s2:   []string{"d", "d", "e"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "duplicates in both",
			s1:   []string{"a", "a", "b", "b", "c", "c"},
			s2:   []string{"a", "b", "b", "c"},
			want: []string{},
		},
		{
			name: "all duplicates in s1 with one unique",
			s1:   []string{"x", "x", "x", "x", "y"},
			s2:   []string{"x"},
			want: []string{"y"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffSlice(tt.s1, tt.s2)
			if !slices.Equal(got, tt.want) {
				t.Errorf("diffSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffSliceOrderPreservation(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want []string
	}{
		{
			name: "preserve order after sorting",
			s1:   []string{"z", "a", "x", "b", "y"},
			s2:   []string{"b", "x"},
			want: []string{"a", "y", "z"},
		},
		{
			name: "order with duplicates",
			s1:   []string{"c", "a", "b", "a", "c", "b"},
			s2:   []string{"b"},
			want: []string{"a", "a", "c", "c"},
		},
		{
			name: "already sorted",
			s1:   []string{"a", "b", "c", "d"},
			s2:   []string{"b", "d"},
			want: []string{"a", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffSlice(tt.s1, tt.s2)
			if !slices.Equal(got, tt.want) {
				t.Errorf("diffSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffSliceEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want []string
	}{
		{
			name: "s1 elements all before s2 elements",
			s1:   []string{"a", "b", "c", "d", "e"},
			s2:   []string{"f", "g", "h"},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "s1 elements all after s2 elements",
			s1:   []string{"f", "g", "h", "i", "j"},
			s2:   []string{"a", "b", "c"},
			want: []string{"f", "g", "h", "i", "j"},
		},
		{
			name: "s1 elements interleaved with s2",
			s1:   []string{"a", "b", "c", "d", "e", "f", "g"},
			s2:   []string{"b", "d", "f"},
			want: []string{"a", "c", "e", "g"},
		},
		{
			name: "all s2 elements at the end of s1",
			s1:   []string{"a", "b", "c", "d", "e"},
			s2:   []string{"c", "d", "e"},
			want: []string{"a", "b"},
		},
		{
			name: "all s2 elements at the beginning of s1",
			s1:   []string{"a", "b", "c", "d", "e"},
			s2:   []string{"a", "b", "c"},
			want: []string{"d", "e"},
		},
		{
			name: "s2 with unsorted values after sorting",
			s1:   []string{"a", "b", "c", "d"},
			s2:   []string{"d", "c", "b"},
			want: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffSlice(tt.s1, tt.s2)
			if !slices.Equal(got, tt.want) {
				t.Errorf("diffSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffSlicePerformance(t *testing.T) {
	// Test with large s1, small s2
	t.Run("Large s1, small s2", func(t *testing.T) {
		s1 := make([]string, 100000)
		for i := 0; i < 100000; i++ {
			s1[i] = fmt.Sprintf("item_%d", i)
		}
		s2 := []string{"item_50000", "item_75000", "item_99999"}

		start := time.Now()
		result := diffSlice(s1, s2)
		elapsed := time.Since(start)

		expected := 100000 - 3
		if len(result) != expected {
			t.Errorf("Expected %d elements, got %d", expected, len(result))
		}

		t.Logf("Large s1, small s2 took: %v", elapsed)
	})

	// Test with small s1, large s2
	t.Run("Small s1, large s2", func(t *testing.T) {
		s1 := []string{"x", "y", "z"}
		s2 := make([]string, 100000)
		for i := 0; i < 100000; i++ {
			s2[i] = fmt.Sprintf("item_%d", i)
		}

		start := time.Now()
		result := diffSlice(s1, s2)
		elapsed := time.Since(start)

		if len(result) != 3 {
			t.Errorf("Expected 3 elements, got %d", len(result))
		}

		t.Logf("Small s1, large s2 took: %v", elapsed)
	})

	// Test with both large
	t.Run("Both large", func(t *testing.T) {
		s1 := make([]string, 50000)
		s2 := make([]string, 50000)

		for i := 0; i < 50000; i++ {
			s1[i] = fmt.Sprintf("item_%d", i)
			if i < 25000 {
				s2[i] = fmt.Sprintf("item_%d", i)
			} else {
				s2[i] = fmt.Sprintf("other_%d", i)
			}
		}

		start := time.Now()
		result := diffSlice(s1, s2)
		elapsed := time.Since(start)

		// First 25000 items are in s2, next 25000 are not
		expected := 25000
		if len(result) != expected {
			t.Errorf("Expected %d elements, got %d", expected, len(result))
		}

		t.Logf("Both large took: %v", elapsed)
	})

	// Test with worst-case scenario - many comparisons
	t.Run("Worst case - many comparisons", func(t *testing.T) {
		s1 := make([]string, 10000)
		s2 := make([]string, 10000)

		for i := 0; i < 10000; i++ {
			s1[i] = fmt.Sprintf("item_%d", i)
			s2[i] = fmt.Sprintf("item_%d", i+5000)
		}

		start := time.Now()
		result := diffSlice(s1, s2)
		elapsed := time.Since(start)

		// First 5000 items not in s2, next 5000 items are in s2
		expected := 5000
		if len(result) != expected {
			t.Errorf("Expected %d elements, got %d", expected, len(result))
		}

		t.Logf("Worst case took: %v", elapsed)
	})
}

func BenchmarkDiffSlice(b *testing.B) {
	s1 := make([]string, 10000)
	s2 := make([]string, 5000)

	for i := 0; i < 10000; i++ {
		s1[i] = fmt.Sprintf("item_%d", i)
	}
	for i := 0; i < 5000; i++ {
		s2[i] = fmt.Sprintf("item_%d", i*2)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diffSlice(s1, s2)
	}
}
