package main

import "fmt"

func main() {
	x := 5
	fmt.Println("hello", x, false)

	// %T -> type
	// %v -> value
	// %b -> binary
	// %e -> cientific notation
	// %f -> float
	// %s -> string
	fmt.Printf("%T\n", x)
	fmt.Printf("%T %T\n", x, x)
	fmt.Printf("%v\n", x)
	fmt.Printf("%b\n", x)

	y := 3.14159
	fmt.Printf("%e\n", y)
	fmt.Printf("%f\n", y)

	z := "hello"
	fmt.Printf("%s\n", z)

	w := 0.15
	fmt.Printf("\"%10.2f%%\n", w)

	s := fmt.Sprintf("%v %v %v", x, x, x)
	fmt.Println(s, s)
}
