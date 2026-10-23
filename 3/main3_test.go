package main

import (
	"maps"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	if v, ok := m.Get("a"); !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v; want 1, true", v, ok)
	}
	if v, ok := m.Get("missing"); ok || v != 0 {
		t.Errorf("Get(missing) = %d, %v; want 0, false", v, ok)
	}

	m.Add("a", 10)
	if v, _ := m.Get("a"); v != 10 {
		t.Errorf("Add must overwrite, got %d", v)
	}
}

func TestZeroValueIsStored(t *testing.T) {
	m := NewStringIntMap()
	m.Add("zero", 0)
	if v, ok := m.Get("zero"); !ok || v != 0 {
		t.Errorf("Get(zero) = %d, %v; want 0, true", v, ok)
	}
	if !m.Exists("zero") {
		t.Error("Exists(zero) = false")
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Remove("a")
	if m.Exists("a") {
		t.Error("key still exists after Remove")
	}
	m.Remove("missing") // не должно паниковать
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	if m.Exists("a") {
		t.Error("Exists on empty map = true")
	}
	m.Add("a", 1)
	if !m.Exists("a") {
		t.Error("Exists(a) = false")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	cp := m.Copy()
	if !maps.Equal(cp, map[string]int{"a": 1, "b": 2}) {
		t.Fatalf("Copy() = %v", cp)
	}

	cp["a"] = 100
	cp["c"] = 3
	if v, _ := m.Get("a"); v != 1 {
		t.Error("changing copy affected original")
	}
	if m.Exists("c") {
		t.Error("adding to copy affected original")
	}

	m.Add("d", 4)
	if _, ok := cp["d"]; ok {
		t.Error("changing original affected copy")
	}
}
