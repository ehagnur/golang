package main

import (
	"fmt"
	"math"
)

//interfaces contain a list of method signitures
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// helper functions
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func area(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func perimeter(g geometry) float64 {
	fmt.Println(g)
	return g.perimeter()
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 3}

	// measure(r)
	// measure(c)

	// area(r)
	// area(c)

	fmt.Println(r.perimeter())
	fmt.Println(c.perimeter())
}
