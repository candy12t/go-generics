package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Min[T cmp.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v2
	}
	return v1
}

func main() {
	fmt.Println(Max(1, 2))
	fmt.Println(Min(0.5, 0.1))
}
