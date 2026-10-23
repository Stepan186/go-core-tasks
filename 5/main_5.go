package main

import "fmt"

func Intersect(a, b []int) (bool, []int) {
	set := make(map[int]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	res := []int{}
	for _, v := range b {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return len(res) > 0, res
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(Intersect(a, b))
}
