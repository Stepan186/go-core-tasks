package main

import "fmt"

func Difference(slice1, slice2 []string) []string {
	exclude := make(map[string]struct{}, len(slice2))
	for _, s := range slice2 {
		exclude[s] = struct{}{}
	}

	res := make([]string, 0, len(slice1))
	for _, s := range slice1 {
		if _, ok := exclude[s]; !ok {
			res = append(res, s)
		}
	}
	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(Difference(slice1, slice2))
}
