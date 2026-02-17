package main

import "fmt"

func main() {
	x := 2
	if x < 1 {
		fmt.Printf("%v < 1\n", x)
	} else if x > 3 {
		fmt.Printf("%v > 3\n", x)
	} else {
		fmt.Printf("1 <= %v <= 3\n", x)
	}

	switch x {
	case 1, 3:
		fmt.Println("1 or 3")
	default:
		fmt.Println("default")
	}

	switch {
	case x >= 1:
		fmt.Printf("%v >= 1\n", x)
		fallthrough
	default:
		fmt.Println("always print this")
	}
}
