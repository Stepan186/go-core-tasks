package main

import (
	"slices"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		s1, s2 []string
		want   []string
	}{
		{
			"example",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
			[]string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{"no overlap", []string{"a", "b"}, []string{"c"}, []string{"a", "b"}},
		{"full overlap", []string{"a", "b"}, []string{"b", "a"}, []string{}},
		{"empty first", nil, []string{"a"}, []string{}},
		{"empty second", []string{"a", "b"}, nil, []string{"a", "b"}},
		{"duplicates kept", []string{"a", "a", "b"}, []string{"b"}, []string{"a", "a"}},
		{"case sensitive", []string{"Apple"}, []string{"apple"}, []string{"Apple"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.s1, tt.s2); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
