package main

import (
	"fmt"
	"math"
)

type geom interface {
	Area() float64
}

type rect struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (c1 circle) Area() float64 {
	return math.Pi * c1.radius * c1.radius
}

func (r1 rect) Area() float64 {
	return r1.length * r1.breadth
}

func main() {
	fmt.Println("example for interface")
	cir := circle{3}
	rec := rect{3.8, 5}

	shapes := []geom{}
	shapes = append(shapes, cir, rec)

	for _, shape := range shapes {
		switch shape.(type) {
		case circle:
			fmt.Println("area of circle is: ", shape.Area())
		case rect:
			fmt.Println("area of rect is: ", shape.Area())
		default:
			break
		}
	}
}
