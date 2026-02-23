package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println(add(-1, 1))
	fmt.Println(add(uint(1), 1))
	fmt.Println(add(3.0, 0.14))

	var mp map[string]int = map[string]int{"a": 1, "b": 2}
	values := getValues(mp)
	fmt.Println(values)

	fmt.Println(mul(3, 5))

	var slice GenericSlice[int] = []int{2, 3, 5, 7}
	slice.Print()

	var strct GenericStruct[int]
	strct.values = 5
	fmt.Println(strct.values)
}

func add[T int | uint | float64](a T, b T) T {
	return a + b
}

func getValues[K comparable, V any](mp map[K]V) []V {
	values := []V{}

	for _, value := range mp {
		values = append(values, value)
	}
	return values
}

type Number interface {
	int | uint | float64
}

func mul[T Number](a T, b T) T {
	return a * b
}

type GenericSlice[T any] []T

func (g GenericSlice[T]) Print() {
	for _, value := range g {
		fmt.Println(value)
	}
}

type GenericStruct[T any] struct {
	values T
}
