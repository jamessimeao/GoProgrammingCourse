package main

import "fmt"

func main() {
	var a int = 2
	var b int = 3
	var s int = add(a, b)
	fmt.Printf("sum = %v\n", s)

	number, str := manyOut()
	fmt.Println(number, str)
	fmt.Println(manyOut())

	fmt.Printf("%T\n", fmt.Println)
	n, _ := callFunc(fmt.Println, "Called fmt.Println")
	fmt.Printf("%v\n", n)

	double := callIntFunc(
		func(x int) int {
			return x + x
		},
		3)
	fmt.Println(double)

	f1 := getFunc("Hello")
	helloWorld := f1(" World")
	fmt.Println(helloWorld)

	fmt.Printf("sum = %v\n", sum(1, 2, 3, 4, 5))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}...))
}

func add(a int, b int) int {
	return a + b
}

func manyOut() (int, string) {
	return 0, "hello"
}

func callFunc(function func(...any) (int, error), input string) (int, error) {
	return function(input)
}

func callIntFunc(function func(int) int, input int) int {
	return function(input)
}

func getFunc(str string) func(string) string {
	return func(str2 string) string {
		return str + str2
	}
}

func sum(nums ...int) (total int) {
	total = 0
	for _, value := range nums {
		total += value
	}
	return
}
