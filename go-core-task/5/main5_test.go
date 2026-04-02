package main

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

// Тест 1: Базовый случай - пересечение есть
func TestIntersectBasic(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{3, 4, 5, 6, 7}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{3, 4, 5}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест 2: Пересечения нет
func TestIntersectNoMatch(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	found, result := intersectSlice(s1, s2)

	if found {
		t.Error("Expected found=false, got true")
	}

	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}
}

// Тест 3: Дубликаты в первом слайсе (должен остаться только первый)
func TestIntersectWithDuplicatesInFirst(t *testing.T) {
	s1 := []int{1, 2, 2, 2, 3, 4, 4}
	s2 := []int{2, 4, 6}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{2, 4}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v (only first of each duplicate), got %v", expected, result)
	}
}

// Тест 4: Дубликаты во втором слайсе
func TestIntersectWithDuplicatesInSecond(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{2, 2, 2, 3, 3, 5}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{2, 3}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест 5: Дубликаты в обоих слайсах
func TestIntersectWithDuplicatesInBoth(t *testing.T) {
	s1 := []int{1, 1, 2, 2, 2, 3, 4, 4}
	s2 := []int{2, 2, 3, 3, 5, 5}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{2, 2, 3}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v (only first of each duplicate from s1), got %v", expected, result)
	}
}

// Тест 6: Сохранение порядка (отсортированный вывод)
func TestIntersectOrderPreservation(t *testing.T) {
	s1 := []int{10, 1, 20, 2, 30, 3}
	s2 := []int{3, 20, 1, 30, 2, 10}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	// Так как функция сортирует, результат должен быть отсортирован
	expected := []int{1, 2, 3, 10, 20, 30}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected sorted %v, got %v", expected, result)
	}
}

// Тест 7: Пустые слайсы
func TestIntersectEmptySlices(t *testing.T) {
	testCases := []struct {
		name string
		s1   []int
		s2   []int
	}{
		{"Both empty", []int{}, []int{}},
		{"First empty", []int{}, []int{1, 2, 3}},
		{"Second empty", []int{1, 2, 3}, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			found, result := intersectSlice(tc.s1, tc.s2)

			if found {
				t.Error("Expected found=false, got true")
			}

			if len(result) != 0 {
				t.Errorf("Expected empty slice, got %v", result)
			}
		})
	}
}

// Тест 8: Один элемент
func TestIntersectSingleElement(t *testing.T) {
	testCases := []struct {
		name     string
		s1       []int
		s2       []int
		expected []int
		found    bool
	}{
		{"Match", []int{5}, []int{5}, []int{5}, true},
		{"No match", []int{5}, []int{3}, []int{}, false},
		{"Match with duplicate in s2", []int{5}, []int{5, 5, 5}, []int{5}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			found, result := intersectSlice(tc.s1, tc.s2)

			if found != tc.found {
				t.Errorf("Expected found=%v, got %v", tc.found, found)
			}

			if !slices.Equal(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// Тест 9: Отрицательные числа и ноль
func TestIntersectNegativeAndZero(t *testing.T) {
	s1 := []int{-5, -3, 0, 2, 4}
	s2 := []int{-3, 0, 5, 7}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{-3, 0}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест 10: Большие числа
func TestIntersectLargeNumbers(t *testing.T) {
	s1 := []int{1000000, 2000000, 3000000}
	s2 := []int{2000000, 4000000, 6000000}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{2000000}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест 11: Один слайс - подмножество другого
func TestIntersectSubset(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{2, 3}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{2, 3}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест 12: Полное совпадение
func TestIntersectCompleteMatch(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{1, 2, 3, 4, 5}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Бенчмарк 1: Маленькие слайсы
func BenchmarkIntersectSmall(b *testing.B) {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := []int{2, 3, 5, 7, 11}

	for b.Loop() {
		intersectSlice(s1, s2)
	}
}

// Бенчмарк 2: Средние слайсы
func BenchmarkIntersectMedium(b *testing.B) {
	s1 := make([]int, 1000)
	s2 := make([]int, 1000)
	for i := range 1000 {
		s1[i] = i * 2
		s2[i] = i * 3
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectSlice(s1, s2)
	}
}

// Бенчмарк 3: Большие слайсы
func BenchmarkIntersectLarge(b *testing.B) {
	s1 := make([]int, 10000)
	s2 := make([]int, 10000)
	for i := range 10000 {
		s1[i] = i
		s2[i] = i + 5000
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectSlice(s1, s2)
	}
}

// Бенчмарк 4: Много дубликатов
func BenchmarkIntersectManyDuplicates(b *testing.B) {
	s1 := make([]int, 10000)
	s2 := make([]int, 10000)
	for i := range 10000 {
		s1[i] = 42
		s2[i] = 42
	}

	for b.Loop() {
		intersectSlice(s1, s2)
	}
}

// Бенчмарк 5: Нет пересечений
func BenchmarkIntersectNoMatches(b *testing.B) {
	s1 := make([]int, 10000)
	s2 := make([]int, 10000)
	for i := range 10000 {
		s1[i] = i
		s2[i] = i + 10000
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectSlice(s1, s2)
	}
}

// Тест производительности с разными размерами
func TestIntersectPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("Size_%d", size), func(t *testing.T) {
			s1 := make([]int, size)
			s2 := make([]int, size)
			for i := range size {
				s1[i] = i
				s2[i] = i + size/2
			}

			start := time.Now()
			intersectSlice(s1, s2)
			elapsed := time.Since(start)

			t.Logf("Size %d took %v", size, elapsed)

			if elapsed > time.Duration(size)*time.Microsecond*10 {
				t.Logf("Warning: Slow performance for size %d: %v", size, elapsed)
			}
		})
	}
}

// Corner case: Очень много дубликатов в начале
func TestIntersectManyDuplicatesAtStart(t *testing.T) {
	s1 := []int{1, 1, 1, 1, 1, 2, 3, 4}
	s2 := []int{1, 5, 6}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{1}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Corner case: Все элементы - дубликаты
func TestIntersectAllDuplicates(t *testing.T) {
	s1 := []int{5, 5, 5, 5, 5}
	s2 := []int{5, 5, 5}

	found, result := intersectSlice(s1, s2)

	if !found {
		t.Error("Expected found=true, got false")
	}

	expected := []int{5, 5, 5}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v (only one 5), got %v", expected, result)
	}
}
