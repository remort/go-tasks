package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGenerator() <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		for range 10 {
			out <- rng.Intn(100)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return out
}

func main() {
	fmt.Println("Генератор случайных чисел (небуферизированный канал)")
	fmt.Println("----------------------------------------")

	ch := randomNumberGenerator()

	for num := range ch {
		fmt.Printf("Получено: %d\n", num)
	}

	fmt.Println("Генерация завершена")
}
