package main

import "fmt"

func Include[T comparable](vs []T, target T) bool {
	for _, v := range vs {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	vs := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Include(vs, 10))
}
