package main

import (
	"sync/atomic"
	"testing"
	"time"
)

// TestNormalUsage - нормальное использование без паники
func TestNormalUsage(t *testing.T) {
	maxConcurrent := 3
	wg := NewWaitGroupSem(maxConcurrent)

	for id := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
		}(id)
	}

	wg.Wait()
}

// TestAddAfterWait - паника при Add после Wait
func TestAddAfterWait(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when Add called after Wait")
		}
	}()

	maxConcurrent := 2
	wg := NewWaitGroupSem(maxConcurrent)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
	}()

	wg.Wait()
	wg.Add(1)
}

// TestAddAfterSemaphoreEmpty - Add после опустошения семафора ДОЛЖЕН паниковать
func TestAddAfterSemaphoreEmpty(t *testing.T) {
	maxConcurrent := 1
	wg := NewWaitGroupSem(maxConcurrent)

	wg.Add(1)
	done1 := make(chan bool)
	go func() {
		defer wg.Done()
		time.Sleep(20 * time.Millisecond)
		done1 <- true
	}()

	<-done1

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when Add called after semaphore empty")
		}
	}()

	wg.Add(1)
}

// TestConcurrentAdd - конкурентные Add
func TestConcurrentAdd(t *testing.T) {
	maxConcurrent := 5
	wg := NewWaitGroupSem(maxConcurrent)

	for range 100 {
		go func() {
			wg.Add(1)
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
		}()
	}

	time.Sleep(100 * time.Millisecond)
	wg.Wait()
}

// TestMaxConcurrentLimit - проверка ограничения параллелизма
func TestMaxConcurrentLimit(t *testing.T) {
	maxConcurrent := 2
	wg := NewWaitGroupSem(maxConcurrent)

	var maxObserved int32 = 0
	var current int32 = 0

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cur := atomic.AddInt32(&current, 1)
			defer atomic.AddInt32(&current, -1)

			for {
				old := atomic.LoadInt32(&maxObserved)
				if cur <= old {
					break
				}
				if atomic.CompareAndSwapInt32(&maxObserved, old, cur) {
					break
				}
			}

			time.Sleep(5 * time.Millisecond)
		}()
	}

	wg.Wait()

	if maxObserved > int32(maxConcurrent) {
		t.Errorf("Max concurrent %d exceeds limit %d", maxObserved, maxConcurrent)
	}
}

// TestRapidDone - таски выполняются быстрее, чем добавляются
func TestRapidDone(t *testing.T) {
	maxConcurrent := 2
	wg := NewWaitGroupSem(maxConcurrent)

	defer func() {
		if r := recover(); r != nil {
			if r != "WaitGroupSem: Add вызван после полного опустошения семафора" {
				t.Errorf("Wrong panic: %v", r)
			}
		}
	}()

	for id := range 5 {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
		}(id)

		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
}

// TestEdgeCaseZeroAdd - нулевой Add
func TestEdgeCaseZeroAdd(t *testing.T) {
	maxConcurrent := 3
	wg := NewWaitGroupSem(maxConcurrent)

	defer func() {
		if r := recover(); r != nil {
			if r != "WaitGroupSem: Add должен быть больше 0" {
				t.Errorf("Wrong panic: %v", r)
			}
		}
	}()

	wg.Add(0)
	wg.Wait()
}

// Бенчмарки
func BenchmarkWaitGroupSem(b *testing.B) {
	maxConcurrent := 10

	for range b.N {
		wg := NewWaitGroupSem(maxConcurrent)
		for range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
			}()
		}
		wg.Wait()
	}
}
