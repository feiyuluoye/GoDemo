package practices

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectry struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectry) area() float64 {
	return r.width * r.height
}

func (r rectry) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func InterfaceMain() {
	r := rectry{width: 2, height: 3}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
