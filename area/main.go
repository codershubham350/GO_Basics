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

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	sq := square{
		sideLength: 5.26,
	}

	tr := triangle{
		base:   7.25,
		height: 8.36,
	}
	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
