package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestTypeOf(t *testing.T) {
	tests := []struct {
		val  any
		want string
	}{
		{42, "int"},
		{052, "int"},
		{0x2A, "int"},
		{3.14, "float64"},
		{"Golang", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}
	for _, tt := range tests {
		if got := TypeOf(tt.val); got != tt.want {
			t.Errorf("TypeOf(%v) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		val  any
		want string
	}{
		{42, "42"},
		{052, "42"},
		{0x2A, "42"},
		{3.14, "3.14"},
		{"Golang", "Golang"},
		{true, "true"},
		{complex64(1 + 2i), "(1+2i)"},
	}
	for _, tt := range tests {
		if got := ToString(tt.val); got != tt.want {
			t.Errorf("ToString(%v) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestJoinValues(t *testing.T) {
	got := JoinValues(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	want := "4242423.14Golangtrue(1+2i)"
	if got != want {
		t.Errorf("JoinValues() = %q, want %q", got, want)
	}
	if got := JoinValues(); got != "" {
		t.Errorf("JoinValues() with no args = %q, want empty", got)
	}
}

func TestAddSalt(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"abcd", "abgo-2024cd"},
		{"abc", "ago-2024bc"},
		{"", "go-2024"},
		{"привет", "приgo-2024вет"},
	}
	for _, tt := range tests {
		if got := string(AddSalt([]rune(tt.in), salt)); got != tt.want {
			t.Errorf("AddSalt(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestAddSaltDoesNotModifyInput(t *testing.T) {
	in := []rune("abcd")
	AddSalt(in, salt)
	if string(in) != "abcd" {
		t.Errorf("input modified: %q", string(in))
	}
}

func TestHashRunes(t *testing.T) {
	runes := []rune("4242423.14Golangtrue(1+2i)")
	sum := sha256.Sum256([]byte("4242423.14Golgo-2024angtrue(1+2i)"))
	want := hex.EncodeToString(sum[:])

	got := HashRunes(runes, salt)
	if got != want {
		t.Errorf("HashRunes() = %s, want %s", got, want)
	}
	if HashRunes(runes, "other") == got {
		t.Error("hash must depend on salt")
	}
}
