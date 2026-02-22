package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		var result int = divide(1, 0)
		fmt.Println(result)
	*/

	/*
		defer defferedFunction(1)
		defer defferedFunction(2)
		defer defferedFunction(3)
		f()
		defer recoverFromError()
		panic("Panic!")
		fmt.Println("1")
		fmt.Println("2")
	*/
	result, err := safeDivide(6, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func divide(a int, b int) int {
	return a / b
}

func defferedFunction(i int) {
	fmt.Printf("defer %v\n", i)
}

func f() {
	defer defferedFunction(4)
	defer defferedFunction(5)
	defer defferedFunction(6)
}

func recoverFromError() {
	r := recover()
	if r != nil {
		fmt.Println("Error occured:")
		fmt.Println(r)
		fmt.Println("But recovered...")
	} else {
		fmt.Println("No error.")
	}
}

func safeDivide(a int, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, errors.New("Division by 0")
	}
}
