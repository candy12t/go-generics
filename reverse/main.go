package main

import "fmt"

func Reverse[T any](vs []T) []T {
	// https://go.dev/wiki/SliceTricks#reversing
	for i := len(vs)/2 - 1; i >= 0; i-- {
		opp := len(vs) - 1 - i
		vs[i], vs[opp] = vs[opp], vs[i]
	}
	return vs
}

func main() {
	vs := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(vs))
}
