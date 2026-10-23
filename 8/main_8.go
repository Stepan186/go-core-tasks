package main

import (
	"fmt"
	"time"
)

type semaphore chan struct{}

func (s semaphore) acquire() { s <- struct{}{} }
func (s semaphore) release() { <-s }

type WaitGroup struct {
	sem     semaphore
	counter int
	done    chan struct{}
}

func NewWaitGroup() *WaitGroup {
	done := make(chan struct{})
	close(done) // счетчик равен нулю — Wait не блокируется
	return &WaitGroup{sem: make(semaphore, 1), done: done}
}

func (wg *WaitGroup) Add(delta int) {
	wg.sem.acquire()
	defer wg.sem.release()

	if wg.counter == 0 && delta > 0 {
		wg.done = make(chan struct{})
	}
	wg.counter += delta
	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}
	if wg.counter == 0 && delta < 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.sem.acquire()
	done := wg.done
	wg.sem.release()
	<-done
}

func main() {
	wg := NewWaitGroup()
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i*10) * time.Millisecond)
			fmt.Println("горутина", i, "завершена")
		}()
	}
	wg.Wait()
	fmt.Println("все горутины завершены")
}
