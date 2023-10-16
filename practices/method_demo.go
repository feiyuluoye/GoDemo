package practices

import "fmt"

type rect struct {
	width, height int
}

// 指针
func (r *rect) area() int {
	return r.width * r.height
}

// value
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func MethodMain() {
	r := rect{width: 10, height: 5}
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())
}
