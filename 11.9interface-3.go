// simple_interface2.go
package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}


func gI(any interface{}) int {
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}
	return 0 // default value
}


func main() {
	var s Simple = Simple{6}
	fmt.Println(gI(&s))
}