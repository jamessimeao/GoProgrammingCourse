package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "1234"
	y, err := strconv.Atoi(x)
	fmt.Println(y, err)

	var z int64
	const base int = 10 // int != int32 or int64
	const bitSize int = 32
	z, err = strconv.ParseInt(x, base, bitSize)
	fmt.Println(z, err)
}
