package main

import "fmt"

// START1 OMIT
type Point struct {
	x int
	y int
}

type Point3D struct {
	Point // HL
	z     int
}

func main() {
	p := Point{42, -10}
	p.plot()
	p.swap()
	p.plot()

	p3 := Point3D{
		Point{42, -10}, // HL
		77,
	}
	p3.plot()
	p3.swap()
	p3.plot()
}
// END1 OMIT

// START2 OMIT
func (p *Point) plot() {
	fmt.Printf("%d, %d\n", p.x, p.y)
}

func (p *Point) swap() {
	x := p.x
	p.x = p.y
	p.y = x
}

func (p3 *Point3D) plot() {
	fmt.Printf("%d, %d, %d\n", p3.x, p3.y, p3.z)
}

// END2 OMIT
