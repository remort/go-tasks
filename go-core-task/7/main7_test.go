package main

import (
	"sort"
	"testing"
)

// Тест 1: Проверка количества элементов
func TestMergeChannelsCount(t *testing.T) {
	chNum := 5
	msgNum := 3
	expectedCount := chNum * msgNum

	merged := mergeChannels(chNum, msgNum)

	count := 0
	for range merged {
		count++
	}

	if count != expectedCount {
		t.Errorf("Expected %d elements, got %d", expectedCount, count)
	}
}

// Тест 2: Проверка, что все числа в диапазоне
func TestMergeChannelsValues(t *testing.T) {
	chNum := 5
	msgNum := 3

	merged := mergeChannels(chNum, msgNum)

	for v := range merged {
		if v < 0 || v >= msgNum {
			t.Errorf("Value %d out of range [0, %d)", v, msgNum)
		}
	}
}

// Тест 3: Проверка, что каждый канал отправил все свои числа
func TestMergeChannelsCompleteness(t *testing.T) {
	chNum := 5
	msgNum := 3

	merged := mergeChannels(chNum, msgNum)

	// Собираем все значения
	values := make([]int, 0)
	for v := range merged {
		values = append(values, v)
	}

	// Проверяем, что каждое число (0..msgNum-1) встречается chNum раз
	expectedCounts := make(map[int]int)
	for i := 0; i < msgNum; i++ {
		expectedCounts[i] = chNum
	}

	actualCounts := make(map[int]int)
	for _, v := range values {
		actualCounts[v]++
	}

	for num, expected := range expectedCounts {
		if actualCounts[num] != expected {
			t.Errorf("Number %d: expected %d times, got %d", num, expected, actualCounts[num])
		}
	}
}

// Тест 4: С различным количеством каналов
func TestMergeChannelsDifferentChannelCounts(t *testing.T) {
	testCases := []struct {
		name     string
		chNum    int
		msgNum   int
		expected int
	}{
		{"1 channel, 5 messages", 1, 5, 5},
		{"3 channels, 10 messages", 3, 10, 30},
		{"10 channels, 1 message", 10, 1, 10},
		{"0 channels, 5 messages", 0, 5, 0},
		{"5 channels, 0 messages", 5, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			merged := mergeChannels(tc.chNum, tc.msgNum)

			count := 0
			for range merged {
				count++
			}

			if count != tc.expected {
				t.Errorf("Expected %d elements, got %d", tc.expected, count)
			}
		})
	}
}

// Тест 5: Проверка закрытия канала
func TestMergeChannelsChannelClosed(t *testing.T) {
	chNum := 3
	msgNum := 2

	merged := mergeChannels(chNum, msgNum)

	// Читаем все данные
	for range merged {
		// потребляем
	}

	// Пытаемся прочитать из закрытого канала
	_, ok := <-merged
	if ok {
		t.Error("Channel should be closed after all data is read")
	}
}

// Тест 6: Проверка порядка (хотя порядок не гарантирован)
func TestMergeChannelsOrder(t *testing.T) {
	chNum := 3
	msgNum := 5

	merged := mergeChannels(chNum, msgNum)

	// Собираем все значения
	values := make([]int, 0)
	for v := range merged {
		values = append(values, v)
	}

	// Сортируем для проверки состава
	sorted := make([]int, len(values))
	copy(sorted, values)
	sort.Ints(sorted)

	// Ожидаем: chNum раз каждое число от 0 до msgNum-1
	expected := make([]int, 0, chNum*msgNum)
	for i := 0; i < msgNum; i++ {
		for j := 0; j < chNum; j++ {
			expected = append(expected, i)
		}
	}

	if len(sorted) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(sorted))
	}

	for i := 0; i < len(sorted); i++ {
		if sorted[i] != expected[i] {
			t.Errorf("At position %d: expected %d, got %d", i, expected[i], sorted[i])
		}
	}
}

// Тест 7: Параллельный запуск (без гонок)
func TestMergeChannelsConcurrent(t *testing.T) {
	chNum := 10
	msgNum := 5

	// Запускаем несколько раз параллельно
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			merged := mergeChannels(chNum, msgNum)
			count := 0
			for range merged {
				count++
			}
			if count != chNum*msgNum {
				t.Errorf("Expected %d, got %d", chNum*msgNum, count)
			}
			done <- true
		}()
	}

	// Ждем завершения всех
	for i := 0; i < 10; i++ {
		<-done
	}
}

// Бенчмарки
func BenchmarkMergeChannelsSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merged := mergeChannels(5, 3)
		for range merged {
			// потребляем
		}
	}
}

func BenchmarkMergeChannelsMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merged := mergeChannels(100, 10)
		for range merged {
			// потребляем
		}
	}
}

func BenchmarkMergeChannelsLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merged := mergeChannels(1000, 100)
		for range merged {
			// потребляем
		}
	}
}

// Тест на утечку горутин
func TestMergeChannelsNoGoroutineLeak(t *testing.T) {
	// Запускаем много раз и проверяем, что все горутины завершаются
	for i := 0; i < 100; i++ {
		merged := mergeChannels(10, 5)
		for range merged {
			// потребляем
		}
	}
	// Если есть утечка горутин, тест может замедлиться или упасть
}

// Тест с очень большим количеством данных
func TestMergeChannelsLargeData(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large data test in short mode")
	}

	chNum := 100
	msgNum := 1000
	expected := chNum * msgNum

	merged := mergeChannels(chNum, msgNum)

	count := 0
	for range merged {
		count++
		if count > expected {
			t.Errorf("Count exceeded expected: %d > %d", count, expected)
			break
		}
	}

	if count != expected {
		t.Errorf("Expected %d elements, got %d", expected, count)
	}
}
