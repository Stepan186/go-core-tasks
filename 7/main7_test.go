package main

import (
	"slices"
	"testing"
	"time"
)

func collect[T any](t *testing.T, ch <-chan T) []T {
	t.Helper()
	var res []T
	timeout := time.After(time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return res
			}
			res = append(res, v)
		case <-timeout:
			t.Fatal("merged channel not closed")
		}
	}
}

func TestMerge(t *testing.T) {
	got := collect(t, Merge(produce(1, 2, 3), produce(10, 20), produce(100)))
	slices.Sort(got)
	want := []int{1, 2, 3, 10, 20, 100}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMergeKeepsOrderWithinChannel(t *testing.T) {
	got := collect(t, Merge(produce(1, 2, 3, 4, 5), produce(-1, -2, -3)))
	var pos, neg []int
	for _, v := range got {
		if v > 0 {
			pos = append(pos, v)
		} else {
			neg = append(neg, v)
		}
	}
	if !slices.Equal(pos, []int{1, 2, 3, 4, 5}) || !slices.Equal(neg, []int{-1, -2, -3}) {
		t.Errorf("order within a channel broken: %v", got)
	}
}

func TestMergeNoChannels(t *testing.T) {
	if got := collect(t, Merge[int]()); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestMergeEmptyChannels(t *testing.T) {
	if got := collect(t, Merge(produce[string](), produce[string]())); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestMergeManyChannels(t *testing.T) {
	const n = 100
	chans := make([]<-chan int, n)
	for i := range n {
		chans[i] = produce(i)
	}
	got := collect(t, Merge(chans...))
	slices.Sort(got)
	for i := range n {
		if got[i] != i {
			t.Fatalf("got %v", got)
		}
	}
}
