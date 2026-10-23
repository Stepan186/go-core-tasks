package main

import "fmt"

// Cube читает числа uint8 из in по мере поступления, преобразует их в float64,
// возводит в куб и пишет результат в out. Закрывает out, когда in закрыт.
func Cube(in <-chan uint8, out chan<- float64) {
	defer close(out)
	for v := range in {
		f := float64(v)
		out <- f * f * f
	}
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		for _, v := range []uint8{0, 1, 2, 3, 10, 255} {
			in <- v
		}
	}()
	go Cube(in, out)

	for v := range out {
		fmt.Println(v)
	}
}
