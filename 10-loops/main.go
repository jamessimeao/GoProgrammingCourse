package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// "while" loop
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}

	// iterate through string
	s := "hello"
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
}
