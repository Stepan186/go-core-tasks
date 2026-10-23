package main

import (
	"slices"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name   string
		a, b   []int
		wantOk bool
		want   []int
	}{
		{"example", []int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, true, []int{64, 3}},
		{"no intersection", []int{1, 2}, []int{3, 4}, false, []int{}},
		{"empty a", nil, []int{1}, false, []int{}},
		{"empty b", []int{1}, nil, false, []int{}},
		{"both empty", nil, nil, false, []int{}},
		{"duplicates", []int{1, 1, 2}, []int{1, 1, 1}, true, []int{1}},
		{"negative and zero", []int{-1, 0, 5}, []int{0, -1}, true, []int{0, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, got := Intersect(tt.a, tt.b)
			if ok != tt.wantOk {
				t.Errorf("ok = %v, want %v", ok, tt.wantOk)
			}
			if got == nil {
				t.Error("result must be empty slice, not nil")
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
