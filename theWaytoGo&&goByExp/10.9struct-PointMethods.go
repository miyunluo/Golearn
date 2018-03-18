package main 

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Scale(s float64) {
	p.x *= s
	p.y *= s
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.x * p.x + p.y * p.y))
}

type Point3 struct {
	x,y,z float64
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y + p.z*p.z))
}

type Polar struct {
	r,t float64
}

func (p Polar) Abs() float64 {
	return p.r
}

func main() {
	p0 := new(Point)
	p0.x = 1
	p0.y = 2
	fmt.Printf("The length of p0 is: %f\n", p0.Abs())

	p1 := &Point{4, 5}
	fmt.Printf("The length of p1 is: %f\n", p1.Abs())
}