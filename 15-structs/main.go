package main

import "fmt"

func main() {
	var person Person = Person{
		10,
		"john"}
	fmt.Println(person)
	person = Person{name: "smith", age: 15}
	fmt.Println(person)
	person.age = 17
	person.name = "mark"
	fmt.Println(person)
	fmt.Println(getName(person))

	var a Person2 = Person2{age: 21}
	a.f = func(x int) int {
		return 5 * x
	}
	fmt.Println(a.age)
	fmt.Println(a.f(5))

	person.setName("name changed")
	fmt.Println(person.getName())
}

type Person struct {
	age  uint
	name string
}

func getName(person Person) string {
	return person.name
}

// Method as a function that is a member of the struct
type Person2 struct {
	name string
	age  int
	f    func(int) int
}

// Better way to make a method
func (person Person) getName() string {
	return person.name
}

// Doesn't work, person is a copy, not a reference
func (person Person) setName(newName string) {
	person.name = newName
}
