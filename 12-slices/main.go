package main

import "fmt"

func main() {
	// slice from array
	a := [5]int{0, 1, 2, 3, 4}
	slice := a[1:3]
	slice[0] = 100
	fmt.Println(a, slice)
	fmt.Println(len(slice), cap(slice))

	// start with a slice
	slice2 := []int{1, 2}
	fmt.Printf("%T\n", slice2)
	fmt.Println(slice2)

	// Appending to a slice works similar to a C++ std::vector
	slice3 := []int{0, 1}
	for i := 2; i < 10; i++ {
		slice3 = append(slice3, i)
		fmt.Println(slice3, len(slice3), cap(slice3))
	}

	// Make a slice with given length and capacity
	slice4 := make([]int, 5, 10)
	fmt.Println(slice4, len(slice4), cap(slice4))

	// Iterate through slice
	slice5 := []int{0, 1, 2, 3, 4, 5}

	for _, value := range slice5 {
		fmt.Println(value)
	}

	// Modify a slice
	modifySlice(slice5)
	fmt.Println(slice5)
}

func modifySlice(slice []int) {
	slice[0] = 55
}
