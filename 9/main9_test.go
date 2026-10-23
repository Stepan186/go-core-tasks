package main

import (
	"slices"
	"testing"
	"time"
)

func runPipeline(t *testing.T, input []uint8) []float64 {
	t.Helper()
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		for _, v := range input {
			in <- v
		}
	}()
	go Cube(in, out)

	var res []float64
	timeout := time.After(time.Second)
	for {
		select {
		case v, ok := <-out:
			if !ok {
				return res
			}
			res = append(res, v)
		case <-timeout:
			t.Fatal("output channel not closed")
		}
	}
}

func TestCube(t *testing.T) {
	got := runPipeline(t, []uint8{0, 1, 2, 3, 10})
	want := []float64{0, 1, 8, 27, 1000}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCubeMaxUint8NoOverflow(t *testing.T) {
	got := runPipeline(t, []uint8{255})
	if len(got) != 1 || got[0] != 16581375 {
		t.Errorf("got %v, want [16581375]", got)
	}
}

func TestCubeEmpty(t *testing.T) {
	if got := runPipeline(t, nil); len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestCubeProcessesAsTheyArrive(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	go Cube(in, out)
	defer close(in)

	for _, v := range []uint8{2, 4} {
		in <- v
		select {
		case got := <-out:
			if want := float64(v) * float64(v) * float64(v); got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		case <-time.After(time.Second):
			t.Fatal("value not processed before input closed")
		}
	}
}
