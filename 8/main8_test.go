package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func waitWithTimeout(t *testing.T, wg *WaitGroup) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return")
	}
}

func TestWaitWithoutAdd(t *testing.T) {
	waitWithTimeout(t, NewWaitGroup())
}

func TestWaitAllGoroutines(t *testing.T) {
	const n = 100
	wg := NewWaitGroup()
	var cnt atomic.Int64
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			cnt.Add(1)
		}()
	}
	waitWithTimeout(t, wg)
	if got := cnt.Load(); got != n {
		t.Errorf("completed %d goroutines, want %d", got, n)
	}
}

func TestWaitBlocks(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)

	returned := make(chan struct{})
	go func() {
		wg.Wait()
		close(returned)
	}()

	select {
	case <-returned:
		t.Fatal("Wait returned before Done")
	case <-time.After(50 * time.Millisecond):
	}

	wg.Done()
	select {
	case <-returned:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after Done")
	}
}

func TestMultipleWaiters(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(2)

	const waiters = 10
	var released atomic.Int64
	allDone := NewWaitGroup()
	allDone.Add(waiters)
	for range waiters {
		go func() {
			defer allDone.Done()
			wg.Wait()
			released.Add(1)
		}()
	}

	wg.Done()
	wg.Done()
	waitWithTimeout(t, allDone)
	if got := released.Load(); got != waiters {
		t.Errorf("released %d waiters, want %d", got, waiters)
	}
}

func TestReuse(t *testing.T) {
	wg := NewWaitGroup()
	for range 3 {
		wg.Add(2)
		go wg.Done()
		go wg.Done()
		waitWithTimeout(t, wg)
	}
}

func TestAddN(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(3)
	for range 3 {
		go wg.Done()
	}
	waitWithTimeout(t, wg)
}

func TestNegativeCounterPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected panic on negative counter")
		}
	}()
	NewWaitGroup().Done()
}
