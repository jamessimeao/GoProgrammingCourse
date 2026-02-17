package main

import (
	"fmt"
	"math"
)

func main() {
	x := int8(7)
	y := 1000
	z := int(x) + y

	fmt.Println(z)

	w := 3.14 / 10
	fmt.Printf("%T\n", w)
	fmt.Println(w)

	s := "hello"
	// Number doesn't convert to string with casting,
	// this gives the ASC character.
	a := 97 // asc code for a is 97, in decimal
	t := s + string(a)
	fmt.Println(t)

	// Proper way
	t = s + fmt.Sprint(a)

	fmt.Println(t)

	fmt.Println(math.Abs(-5))
}
