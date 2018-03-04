package main 

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

func (car *Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car
}

func (mer *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi!")
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

func main() {
	m := Mercedes{Car{4, nil}}
	fmt.Println("A Mercedes has", m.numberOfWheels(), "wheels.")
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
