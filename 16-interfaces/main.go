package main

import (
	"fmt"
	"math"
)

func main() {
	var rectangle Shape2D = Rectangle{width: 2, height: 3}
	var circle Shape2D = Circle{radius: 1}
	var shapes []Shape2D = []Shape2D{rectangle, circle}
	for i := 0; i < len(shapes); i++ {
		fmt.Printf("perimeter = %v\n", shapes[i].perimeter())
		fmt.Printf("area = %v\n", shapes[i].area())
	}
}

type Shape2D interface {
	perimeter() float32
	area() float32
}

type Rectangle struct {
	width  float32
	height float32
}

func (rectangle Rectangle) perimeter() float32 {
	return 2 * (rectangle.height + rectangle.width)
}

func (rectangle Rectangle) area() float32 {
	return rectangle.height * rectangle.width
}

type Circle struct {
	radius float32
}

func (circle Circle) perimeter() float32 {
	return 2 * math.Pi * circle.radius
}

func (circle Circle) area() float32 {
	return math.Pi * circle.radius * circle.radius
}
