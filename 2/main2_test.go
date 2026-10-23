package main

import (
	"errors"
	"slices"
	"testing"
)

func TestRandomSlice(t *testing.T) {
	s := randomSlice(10, 100)
	if len(s) != 10 {
		t.Fatalf("len = %d, want 10", len(s))
	}
	for _, v := range s {
		if v < 0 || v >= 100 {
			t.Errorf("value %d out of range [0, 100)", v)
		}
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"mixed", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8, 10}},
		{"negative", []int{-3, -2, 0, 1}, []int{-2, 0}},
		{"no even", []int{1, 3, 5}, []int{}},
		{"empty", nil, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceExample(tt.in); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	orig := make([]int, 3, 10)
	copy(orig, []int{1, 2, 3})

	got := addElements(orig, 4)
	if !slices.Equal(got, []int{1, 2, 3, 4}) {
		t.Fatalf("got %v", got)
	}

	got[0] = 100
	if orig[0] != 1 {
		t.Error("result shares memory with original")
	}
	if !slices.Equal(addElements(nil, 5), []int{5}) {
		t.Error("addElements(nil, 5) failed")
	}
}

func TestCopySlice(t *testing.T) {
	orig := []int{1, 2, 3}
	cp := copySlice(orig)
	if !slices.Equal(cp, orig) {
		t.Fatalf("got %v, want %v", cp, orig)
	}

	orig[0] = 100
	if cp[0] != 1 {
		t.Error("changes in original affected copy")
	}
	if got := copySlice(nil); len(got) != 0 {
		t.Errorf("copySlice(nil) = %v", got)
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		idx  int
		want []int
	}{
		{"first", 0, []int{2, 3, 4}},
		{"middle", 2, []int{1, 2, 4}},
		{"last", 3, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := []int{1, 2, 3, 4}
			got, err := removeElement(orig, tt.idx)
			if err != nil {
				t.Fatal(err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if !slices.Equal(orig, []int{1, 2, 3, 4}) {
				t.Errorf("original modified: %v", orig)
			}
		})
	}
}

func TestRemoveElementOutOfRange(t *testing.T) {
	for _, idx := range []int{-1, 4, 100} {
		if _, err := removeElement([]int{1, 2, 3, 4}, idx); !errors.Is(err, ErrIndexOutOfRange) {
			t.Errorf("idx %d: err = %v, want ErrIndexOutOfRange", idx, err)
		}
	}
}
