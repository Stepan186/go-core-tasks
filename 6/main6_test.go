package main

import (
	"context"
	"testing"
	"time"
)

func TestRandomGeneratorRange(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gen := RandomGenerator(ctx, 10)
	for range 1000 {
		if v := <-gen; v < 0 || v >= 10 {
			t.Fatalf("value %d out of range [0, 10)", v)
		}
	}
}

func TestRandomGeneratorUnbuffered(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if c := cap(RandomGenerator(ctx, 10)); c != 0 {
		t.Errorf("cap = %d, want 0 (unbuffered)", c)
	}
}

func TestRandomGeneratorIsRandom(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gen := RandomGenerator(ctx, 1_000_000)
	seen := make(map[int]struct{})
	for range 100 {
		seen[<-gen] = struct{}{}
	}
	if len(seen) < 2 {
		t.Error("generator returns the same value")
	}
}

func TestRandomGeneratorClosesOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	gen := RandomGenerator(ctx, 10)
	<-gen
	cancel()

	timeout := time.After(time.Second)
	for {
		select {
		case _, ok := <-gen:
			if !ok {
				return
			}
		case <-timeout:
			t.Fatal("channel not closed after cancel")
		}
	}
}
