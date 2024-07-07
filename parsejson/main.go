package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ParseJSON[T any](data []byte) (*T, error) {
	v := new(T)
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func main() {
	input := `{"name": "Mike", "age": 20}`
	v, err := ParseJSON[User]([]byte(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", v)
}
