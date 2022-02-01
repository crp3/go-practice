package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{15.0, 4.0}
	s := square{10.0}

	printArea(t)
	printArea(s)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) * 0.5
}
