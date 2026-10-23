package main

import (
	"context"
	"fmt"
	"math/rand/v2"
)

func RandomGenerator(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- rand.IntN(max):
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gen := RandomGenerator(ctx, 100)
	for range 5 {
		fmt.Println(<-gen)
	}
}
