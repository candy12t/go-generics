package main

import "fmt"

type Tuple[T1, T2 any] struct {
	V1 T1
	V2 T2
}

func New[T1, T2 any](v1 T1, v2 T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{
		V1: v1,
		V2: v2,
	}
}

func main() {
	t := New("hoge", 1)
	fmt.Println(t.V1, t.V2)
}
