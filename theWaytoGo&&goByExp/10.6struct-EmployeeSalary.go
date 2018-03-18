package main

import "fmt"

type emp struct {
	salary float32
}

func (this *emp) giveRaise(percent float32){
	this.salary += this.salary * percent
}

func main() {
	e := new(emp)
	e.salary = 1000
	e.giveRaise(0.02)
	fmt.Printf("Employee earns %f now.\n", e.salary)	
}