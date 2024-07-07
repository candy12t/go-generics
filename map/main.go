package main

import "fmt"

func Map[T1, T2 any](vs []T1, f func(T1) T2) []T2 {
	result := make([]T2, 0, len(vs))
	for _, v := range vs {
		result = append(result, f(v))
	}
	return result
}

func main() {
	vs := []int{1, 2, 3, 4, 5, 6}
	result := Map(vs, func(v int) string {
		return fmt.Sprintf("#%d", v)
	})
	fmt.Println(result) // [#1 #2 #3 #4 #5 #6]
}
