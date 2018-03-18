package main 

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	n int
}

func (s *Simple) Get() int {
	return s.n
}

func (s*Simple) Set(i int) {
	s.n = i
}

func main() {
	s := &Simple{1}
	var ser Simpler
	ser = s
	fmt.Println("Get: ", ser.Get())
}