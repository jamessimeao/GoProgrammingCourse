package main

import "fmt"

func main() {
	// signed integers
	var x1 int8 = -1
	var x2 int16 = -2
	var x3 int32 = -3
	var x4 int64 = -4
	var r rune = -5 // alias for int32
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(r)

	// unsigned integers
	var x5 uint8 = 1
	var x6 uint16 = 2
	var x7 uint32 = 3
	var x8 uint64 = 4
	fmt.Println(x5)
	fmt.Println(x6)
	fmt.Println(x7)
	fmt.Println(x8)

	// floats
	const pi float32 = 3.1415
	const e float64 = 2.71828
	fmt.Println(pi)
	fmt.Println(e)

	// string
	const s string = "Hello"
	fmt.Println(s)
}
