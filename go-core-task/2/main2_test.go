package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "mixed numbers",
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{2, 4, 6, 8, 10},
		},
		{
			name:  "all even",
			slice: []int{2, 4, 6, 8},
			want:  []int{2, 4, 6, 8},
		},
		{
			name:  "all odd",
			slice: []int{1, 3, 5, 7},
			want:  []int{},
		},
		{
			name:  "empty slice",
			slice: []int{},
			want:  []int{},
		},
		{
			name:  "single even",
			slice: []int{2},
			want:  []int{2},
		},
		{
			name:  "single odd",
			slice: []int{1},
			want:  []int{},
		},
		{
			name:  "with zero",
			slice: []int{0, 1, 2, 3},
			want:  []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		element int
		want    []int
	}{
		{
			name:    "append to non-empty slice",
			slice:   []int{10, 20, 30},
			element: 40,
			want:    []int{10, 20, 30, 40},
		},
		{
			name:    "append to empty slice",
			slice:   []int{},
			element: 10,
			want:    []int{10},
		},
		{
			name:    "append to single element slice",
			slice:   []int{5},
			element: 10,
			want:    []int{5, 10},
		},
		{
			name:    "append negative number",
			slice:   []int{1, 2, 3},
			element: -5,
			want:    []int{1, 2, 3, -5},
		},
		{
			name:    "append zero",
			slice:   []int{1, 2, 3},
			element: 0,
			want:    []int{1, 2, 3, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.slice, tt.element)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddElementsLowLevel(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		element int
		want    []int
	}{
		{
			name:    "append to non-empty slice",
			slice:   []int{10, 20, 30},
			element: 40,
			want:    []int{10, 20, 30, 40},
		},
		{
			name:    "append to empty slice",
			slice:   []int{},
			element: 10,
			want:    []int{10},
		},
		{
			name:    "append to single element slice",
			slice:   []int{5},
			element: 10,
			want:    []int{5, 10},
		},
		{
			name:    "append negative number",
			slice:   []int{1, 2, 3},
			element: -5,
			want:    []int{1, 2, 3, -5},
		},
		{
			name:    "append zero",
			slice:   []int{1, 2, 3},
			element: 0,
			want:    []int{1, 2, 3, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElementsLowLevel(tt.slice, tt.element)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElementsLowLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		slice interface{}
	}{
		{
			name:  "integers",
			slice: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "strings",
			slice: []string{"apple", "banana", "cherry"},
		},
		{
			name:  "bytes",
			slice: []byte{'r', 'o', 'a', 'd'},
		},
		{
			name:  "empty slice",
			slice: []string{},
		},
		{
			name:  "single element",
			slice: []int{42},
		},
		{
			name:  "bool slice",
			slice: []bool{true, false, true},
		},
		{
			name:  "float64 slice",
			slice: []float64{1.1, 2.2, 3.3},
		},
		{
			name:  "interface slice",
			slice: []interface{}{1, "two", 3.0, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use type assertion to get the actual slice
			switch v := tt.slice.(type) {
			case []int:
				original := v
				copied := copySlice(original)
				if !reflect.DeepEqual(original, copied) {
					t.Errorf("copySlice() = %v, want %v", copied, original)
				}
				// Verify deep copy
				if len(copied) > 0 {
					copied[0] = copied[0] + 1
					if reflect.DeepEqual(original, copied) {
						t.Errorf("copySlice() returned reference, not a deep copy")
					}
				}
			case []string:
				original := v
				copied := copySlice(original)
				if !reflect.DeepEqual(original, copied) {
					t.Errorf("copySlice() = %v, want %v", copied, original)
				}
			case []byte:
				original := v
				copied := copySlice(original)
				if !reflect.DeepEqual(original, copied) {
					t.Errorf("copySlice() = %v, want %v", copied, original)
				}
			case []bool:
				original := v
				copied := copySlice(original)
				if !reflect.DeepEqual(original, copied) {
					t.Errorf("copySlice() = %v, want %v", copied, original)
				}
			case []float64:
				original := v
				copied := copySlice(original)
				if !reflect.DeepEqual(original, copied) {
					t.Errorf("copySlice() = %v, want %v", copied, original)
				}
			case []interface{}:
				original := v
				copied := copySlice(original)
				if !reflect.DeepEqual(original, copied) {
					t.Errorf("copySlice() = %v, want %v", copied, original)
				}
				// Verify deep copy for interface slice
				if len(copied) > 0 {
					copied[0] = "modified"
					if reflect.DeepEqual(original, copied) {
						t.Errorf("copySlice() returned reference, not a deep copy")
					}
				}
			default:
				t.Errorf("unsupported slice type: %T", tt.slice)
			}
		})
	}
}

func TestCopySliceDeepCopy(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copied := copySlice(original)

	// Modify the copy
	copied[0] = 100

	// Original should remain unchanged
	if reflect.DeepEqual(original, copied) {
		t.Errorf("copySlice() did not create a deep copy: original %v, copied %v", original, copied)
	}

	// Verify elements are equal before modification
	original = []int{1, 2, 3, 4, 5}
	copied = copySlice(original)
	if !reflect.DeepEqual(original, copied) {
		t.Errorf("copySlice() = %v, want %v", copied, original)
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		index     int
		want      []string
		wantPanic bool
	}{
		{
			name:  "remove middle element",
			slice: []string{"apple", "pear", "potato", "banana"},
			index: 1,
			want:  []string{"apple", "potato", "banana"},
		},
		{
			name:  "remove first element",
			slice: []string{"apple", "pear", "potato"},
			index: 0,
			want:  []string{"pear", "potato"},
		},
		{
			name:  "remove last element",
			slice: []string{"apple", "pear", "potato"},
			index: 2,
			want:  []string{"apple", "pear"},
		},
		{
			name:  "remove from single element slice",
			slice: []string{"apple"},
			index: 0,
			want:  []string{},
		},
		{
			name:  "invalid index negative",
			slice: []string{"apple", "pear", "potato"},
			index: -1,
			want:  []string{"apple", "pear", "potato"},
		},
		{
			name:  "invalid index out of range",
			slice: []string{"apple", "pear", "potato"},
			index: 10,
			want:  []string{"apple", "pear", "potato"},
		},
		{
			name:  "empty slice with invalid index",
			slice: []string{},
			index: 0,
			want:  []string{},
		},
		{
			name:  "remove with index equals length",
			slice: []string{"apple", "pear", "potato"},
			index: 3,
			want:  []string{"apple", "pear", "potato"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.slice, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElementWithInts(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		index int
		want  []int
	}{
		{
			name:  "remove middle element",
			slice: []int{10, 20, 30, 40, 50},
			index: 2,
			want:  []int{10, 20, 40, 50},
		},
		{
			name:  "remove first element",
			slice: []int{10, 20, 30},
			index: 0,
			want:  []int{20, 30},
		},
		{
			name:  "remove last element",
			slice: []int{10, 20, 30},
			index: 2,
			want:  []int{10, 20},
		},
		{
			name:  "invalid index returns copy",
			slice: []int{10, 20, 30},
			index: 5,
			want:  []int{10, 20, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.slice, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElementDoesNotModifyOriginal(t *testing.T) {
	original := []string{"apple", "pear", "potato"}
	removed := removeElement(original, 1)

	// Original should remain unchanged
	expectedOriginal := []string{"apple", "pear", "potato"}
	if !reflect.DeepEqual(original, expectedOriginal) {
		t.Errorf("Original slice was modified: got %v, want %v", original, expectedOriginal)
	}

	// Removed slice should be correct
	expectedRemoved := []string{"apple", "potato"}
	if !reflect.DeepEqual(removed, expectedRemoved) {
		t.Errorf("removeElement() = %v, want %v", removed, expectedRemoved)
	}
}

func TestAddElementsDoesNotModifyOriginal(t *testing.T) {
	original := []int{10, 20, 30}
	added := addElements(original, 40)

	if !reflect.DeepEqual(original, []int{10, 20, 30}) {
		t.Errorf("Original slice was modified: got %v, want %v", original, []int{10, 20, 30})
	}

	expectedAdded := []int{10, 20, 30, 40}
	if !reflect.DeepEqual(added, expectedAdded) {
		t.Errorf("addElements() = %v, want %v", added, expectedAdded)
	}
}

func TestAddElementsLowLevelDoesNotModifyOriginal(t *testing.T) {
	original := []int{10, 20, 30}
	added := addElementsLowLevel(original, 40)

	if !reflect.DeepEqual(original, []int{10, 20, 30}) {
		t.Errorf("Original slice was modified: got %v, want %v", original, []int{10, 20, 30})
	}

	expectedAdded := []int{10, 20, 30, 40}
	if !reflect.DeepEqual(added, expectedAdded) {
		t.Errorf("addElementsLowLevel() = %v, want %v", added, expectedAdded)
	}
}
