package main

import "fmt"

type area interface {
	getArea() float64
}

type shape interface {
	printArea(a area)
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	s := square{sideLength: 10}
	printArea(s)

	t := triangle{base: 5, height: 10}
	printArea(t)
}

func printArea(a area) {
	fmt.Println(a.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
