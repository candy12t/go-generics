package main

import "fmt"

func Filter[T any](vs []T, f func(T) bool) []T {
	result := make([]T, 0, len(vs))
	for _, v := range vs {
		if !f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func main() {
	vs := []int{1, 2, 3, 4, 5, 6}
	result := Filter(vs, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(result) // [2 4 6]
}
