package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type WaitGroupSem struct {
	sem     chan struct{} // буферизированный семафор
	done    chan struct{} // сигнал завершения
	started int32
}

func NewWaitGroupSem(maxConcurrent int) *WaitGroupSem {
	return &WaitGroupSem{
		sem:  make(chan struct{}, maxConcurrent),
		done: make(chan struct{}),
	}
}

func (wg *WaitGroupSem) Add(delta int) {
	if delta <= 0 {
		panic("WaitGroupSem: Add должен быть больше 0")
	}
	if atomic.LoadInt32(&wg.started) != 0 && len(wg.sem) == 0 {
		panic("WaitGroupSem: Add вызван после полного опустошения семафора")
	}

	for range delta {
		fmt.Println("Ожидаю семафор...")
		wg.sem <- struct{}{} // захват (блокируется при maxConcurrent)
		atomic.StoreInt32(&wg.started, 1)
		fmt.Println("Захватил семафор...")
	}
}

func (wg *WaitGroupSem) Done() {
	<-wg.sem // освобождаем

	// Если семафор пуст - все горутины завершились
	// Тут возможна потенциальная гонка если таски будут завершаться быстрее чем добавляться новые.
	// Но кажется waitGroup не должна использоваться после того как текущие таски в ней выполнились
	// Так что этим кейсом я пренебрегу здесь. Иначе надо наворачивать много синх. примитивов
	// типа count, is_started и работать с ними. Что сильно увеличивает когнитивную сложность.
	if len(wg.sem) == 0 {
		fmt.Println("Закрываем канал wg...")
		close(wg.done)
	}
}

func (wg *WaitGroupSem) Wait() {
	fmt.Println("Wait ждет...")
	<-wg.done
	fmt.Println("Wait дождался...")
}

func main() {
	maxConcurrent := 3
	wg := NewWaitGroupSem(maxConcurrent)

	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d началась (активных: %d)\n", id, len(wg.sem))
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Горутина %d завершилась (активных: %d)\n", id, len(wg.sem)-1)
		}(i)
	}

	fmt.Println("Ожидаем завершения...")
	wg.Wait()
	fmt.Println("Все готово!")
}
