package main

import "fmt"

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimeter() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}

type Triangle struct {
	base   float32
	height float32
}

func (tr *Triangle) Area() float32 {
	return 0.5 * tr.base * tr.height
}


func main() {
	var areaIntf AreaInterface
	var periIntf PeriInterface

	sq := new(Square)
	sq.side = 5
	tr := new(Triangle)
	tr.base = 3
	tr.height = 5

	areaIntf = sq
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	periIntf = sq
	fmt.Printf("The square has perimeter: %f\n", periIntf.Perimeter())

	areaIntf = tr
	fmt.Printf("The triangle has area: %f\n", areaIntf.Area())
}
