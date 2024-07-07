package main

import (
	"fmt"
)

func Uniq[T comparable](vs []T) []T {
	result := make([]T, 0, len(vs))
	flags := make(map[T]bool, len(vs))
	for _, v := range vs {
		if _, ok := flags[v]; !ok {
			result = append(result, v)
			flags[v] = true
		}
	}
	return result
}

func main() {
	vs := []int{1, 2, 3, 1, 3, 5, 5}
	fmt.Println(Uniq(vs))
}
