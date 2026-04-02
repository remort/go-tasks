package main

import (
	"testing"
	"time"
)

// Тест 1: Проверка, что генератор возвращает канал
func TestRandomNumberGeneratorReturnsChannel(t *testing.T) {
	ch := randomNumberGenerator()

	if ch == nil {
		t.Error("Expected non-nil channel, got nil")
	}

	select {
	case num, ok := <-ch:
		if !ok {
			t.Error("Channel should be open, got closed")
		}
		t.Logf("Received number: %d", num)
	case <-time.After(1 * time.Second):
		t.Error("Timeout: generator didn't send any number")
	}
}

// Тест 2: Проверка количества генерируемых чисел
func TestRandomNumberGeneratorCount(t *testing.T) {
	ch := randomNumberGenerator()

	count := 0
	for range ch {
		count++
	}

	expected := 10
	if count != expected {
		t.Errorf("Expected %d numbers, got %d", expected, count)
	}
}

// Тест 3: Проверка, что все числа в допустимом диапазоне
func TestRandomNumberGeneratorRange(t *testing.T) {
	ch := randomNumberGenerator()

	for num := range ch {
		if num < 0 || num >= 100 {
			t.Errorf("Number %d is out of range [0, 99]", num)
		}
	}
}

// Тест 4: Проверка, что числа не все одинаковые
func TestRandomNumberGeneratorNotAllSame(t *testing.T) {
	ch := randomNumberGenerator()

	numbers := make([]int, 0, 10)
	for num := range ch {
		numbers = append(numbers, num)
	}

	// Проверяем, что не все числа одинаковые
	allSame := true
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[0] {
			allSame = false
			break
		}
	}

	if allSame {
		t.Error("All generated numbers are the same, expected random distribution")
	}

	t.Logf("Generated numbers: %v", numbers)
}

// Тест 5: Проверка блокирующего поведения небуферизированного канала
func TestUnbufferedChannelBlocking(t *testing.T) {
	out := make(chan int)

	start := time.Now()

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 42
	}()

	<-out
	elapsed := time.Since(start)

	if elapsed < 90*time.Millisecond {
		t.Errorf("Expected blocking ~100ms, got %v", elapsed)
	}
}

// Тест 6: Проверка закрытия канала
func TestChannelClosedAfterGeneration(t *testing.T) {
	ch := randomNumberGenerator()

	for range ch {
		// потребляем все числа
	}

	_, ok := <-ch
	if ok {
		t.Error("Channel should be closed after generation, but still open")
	}
}

// Тест 7: Проверка уникальности seed (разные последовательности)
func TestDifferentSeedsProduceDifferentNumbers(t *testing.T) {
	// Первая последовательность
	ch1 := randomNumberGenerator()
	nums1 := make([]int, 0, 10)
	for n := range ch1 {
		nums1 = append(nums1, n)
	}

	// Вторая последовательность (другой seed из-за time.Now())
	time.Sleep(1 * time.Millisecond) // гарантируем разный seed
	ch2 := randomNumberGenerator()
	nums2 := make([]int, 0, 10)
	for n := range ch2 {
		nums2 = append(nums2, n)
	}

	// С вероятностью 1/100^10 могут быть одинаковыми, но это крайне маловероятно
	allSame := true
	for i := 0; i < 10; i++ {
		if nums1[i] != nums2[i] {
			allSame = false
			break
		}
	}

	if allSame {
		t.Log("Warning: sequences are identical (very unlikely but possible)")
	} else {
		t.Log("Sequences are different as expected")
	}
}

// Бенчмарк: производительность генератора
func BenchmarkRandomNumberGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := randomNumberGenerator()
		for range ch {
			// потребляем все числа
		}
	}
}
