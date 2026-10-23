package main

import (
	"fmt"
	"sync"
)

func Merge[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan T) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func produce[T any](values ...T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch
}

func main() {
	a := produce(1, 2, 3)
	b := produce(10, 20)
	c := produce(100, 200, 300, 400)

	for v := range Merge(a, b, c) {
		fmt.Println(v)
	}
}
