package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	res := make(map[string]int, len(m.data))
	for k, v := range m.data {
		res[k] = v
	}
	return res
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.data[key]
	return v, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)
	fmt.Println("Copy:", m.Copy())

	fmt.Println("Exists(one):", m.Exists("one"))
	v, ok := m.Get("two")
	fmt.Println("Get(two):", v, ok)

	m.Remove("one")
	fmt.Println("После Remove(one):", m.Copy())
	v, ok = m.Get("one")
	fmt.Println("Get(one):", v, ok)
}
