package main

import "fmt"

type Number interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~float64 | ~float32
}

func Sum[T Number](v1, v2 T) T {
	return v1 + v2
}

type ID int

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(ID(1), ID(2)))
}
