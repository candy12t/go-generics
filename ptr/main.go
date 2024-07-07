package main

import "fmt"

func Ptr[T any](v T) *T {
	return &v
}

func main() {
	fmt.Println(Ptr("hoge"))
	fmt.Println(Ptr(1))
	fmt.Println(Ptr(true))
}
