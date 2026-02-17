package main

import "fmt"

func main() {
	a := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%T\n", a)

	a[0] = [2]int{7, 8}

	// This function call doesn't modify the original array
	arrayIsCopiedToFunction(a)

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(a[i][j])
		}
	}

	for i, array1d := range a {
		fmt.Println(i, array1d)
		for _, value := range array1d {
			fmt.Println(value)
		}
	}
}

func arrayIsCopiedToFunction(a [3][2]int) {
	// This modifies a copy of the array
	a[0] = [2]int{0, 1}
}
