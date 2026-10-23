package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

var ErrIndexOutOfRange = errors.New("index out of range")

func randomSlice(n, max int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.IntN(max)
	}
	return s
}

func sliceExample(s []int) []int {
	res := make([]int, 0, len(s))
	for _, v := range s {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func addElements(s []int, n int) []int {
	res := make([]int, len(s), len(s)+1)
	copy(res, s)
	return append(res, n)
}

func copySlice(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}

func removeElement(s []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(s) {
		return nil, fmt.Errorf("%w: %d (len %d)", ErrIndexOutOfRange, idx, len(s))
	}
	res := make([]int, 0, len(s)-1)
	res = append(res, s[:idx]...)
	return append(res, s[idx+1:]...), nil
}

func main() {
	originalSlice := randomSlice(10, 100)
	fmt.Println("originalSlice:", originalSlice)

	fmt.Println("sliceExample:", sliceExample(originalSlice))
	fmt.Println("addElements(…, 777):", addElements(originalSlice, 777))

	copied := copySlice(originalSlice)
	originalSlice[0] = -1
	fmt.Println("original после изменения:", originalSlice)
	fmt.Println("copySlice:", copied)

	removed, err := removeElement(originalSlice, 3)
	if err != nil {
		fmt.Println("removeElement:", err)
	} else {
		fmt.Println("removeElement(…, 3):", removed)
	}

	if _, err := removeElement(originalSlice, 42); err != nil {
		fmt.Println("removeElement(…, 42):", err)
	}

	fmt.Println("original в конце:", originalSlice)
}
