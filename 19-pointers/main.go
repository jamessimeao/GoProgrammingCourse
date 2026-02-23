package main

import "fmt"

func main() {
	var x int = 5
	var y = &x
	fmt.Printf("%T %v %v\n", y, y, *y)
	*y = 7
	fmt.Println(x)

	change(y)
	fmt.Println(x)

	var person Person
	person.name = "john"
	fmt.Println(person)

	var personPtr *Person = &person
	personPtr.SetName("smith")
	fmt.Println(person)
	personPtr.SetName2("mark")
	fmt.Println(person)

	// Go has automatic dereference for structs
	person.SetName("smith")
	fmt.Println(person)
	person.SetName2("mark")
	fmt.Println(person)
}

func change(x *int) {
	*x = 100
}

type Person struct {
	name string
}

// Setter
func (person *Person) SetName(newName string) {
	(*person).name = newName // like a C/C++ pointer
}

func (person *Person) SetName2(newName string) {
	person.name = newName // like a C++ reference
}
